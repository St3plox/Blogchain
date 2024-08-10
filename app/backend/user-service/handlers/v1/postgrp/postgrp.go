package postgrp

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"github.com/St3plox/Blogchain/business/core/media"
	"github.com/St3plox/Blogchain/business/core/post"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/gorilla/mux"
)

type Handler struct {
	post  *post.Core
	auth  *auth.Auth
	user  *user.Core
	media *media.Core
}

func New(postCore *post.Core, auth *auth.Auth, userCore *user.Core, mediaCore *media.Core) *Handler {
	return &Handler{
		post:  postCore,
		auth:  auth,
		user:  userCore,
		media: mediaCore,
	}
}

// @Summary Create a new post
// @Description Create a new post with title and content
// @Tags posts
// @Accept json
// @Produce json
// @Param newPost body post.NewPost true "New Post"
// @Success 201 {object} post.Post
// @Failure 400 {object} v1.ErrorResponse
// @Router /posts [post]
func (h *Handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var np post.NewPost
	err := json.NewDecoder(r.Body).Decode(&np)
	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusInternalServerError)
	}
	claims := auth.GetClaims(ctx)

	//TODO: move to postCore
	usr, err := h.user.QueryByID(ctx, claims.Subject)
	if err != nil {
		return v1.NewRequestError(errors.New("user error "+err.Error()), http.StatusNotFound)
	}
	post, err := h.post.Create(ctx, np, usr.AddressHex)
	if err != nil {
		return v1.NewRequestError(errors.New("create error "+err.Error()), http.StatusInternalServerError)
	}

	err = web.Respond(ctx, w, post, http.StatusCreated)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Get posts of user who made the request
// @Description Get all posts by a specific user address
// @Tags posts
// @Produce json
// @Param page query int false "Page number" default(0)
// @Param pageSize query int false "Page size" default(100)
// @Success 200 {array} post.Post
// @Failure 400 {object} v1.ErrorResponse
// @Router /posts [get]
func (h *Handler) GetUserPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := auth.GetClaims(ctx)

	page, pageSize, err := extractPaginationParams(r)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	usr, err := h.user.QueryByID(ctx, claims.Subject)
	if err != nil {
		return v1.NewRequestError(errors.New("user error "+err.Error()), http.StatusNotFound)
	}

	posts, err := h.post.QueryByAddress(ctx, usr.AddressHex, page, pageSize)
	if err != nil {
		return v1.NewRequestError(errors.New("query error "+err.Error()), http.StatusInternalServerError)
	}

	if posts == nil {
		return v1.NewRequestError(errors.New("query error"), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, posts, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Get posts by user address
// @Description Get all posts by a specific user address
// @Tags posts
// @Produce json
// @Param address path string true "User Address"
// @Param page query int false "Page number" default(0)
// @Param pageSize query int false "Page size" default(100)
// @Success 200 {array} post.Post
// @Failure 400 {object} v1.ErrorResponse
// @Router /posts/{address} [get]
func (h *Handler) GetPostsByUserAddress(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	page, pageSize, err := extractPaginationParams(r)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	posts, err := h.post.QueryByAddress(ctx, params["address"], page, pageSize)
	if err != nil {
		return v1.NewRequestError(errors.New("post error "+err.Error()), http.StatusInternalServerError)
	}

	if posts == nil {
		return v1.NewRequestError(errors.New("post error"), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, posts, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Get a post by user address and index
// @Description Get a specific post by user address and index
// @Tags posts
// @Produce json
// @Param address path string true "User Address"
// @Param index path int true "Post Index"
// @Success 200 {object} post.Post
// @Failure 400 {object} v1.ErrorResponse
// @Failure 404 {object} v1.ErrorResponse
// @Router /posts/{address}/{index} [get]
func (h *Handler) GetPostsByUserAddressAndIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	index, err := strconv.ParseUint(params["index"], 10, 64)
	if err != nil {
		return v1.NewRequestError(errors.New("parse error "+err.Error()), http.StatusInternalServerError)
	}

	post, err := h.post.QueryByIndex(ctx, params["address"], index)
	if err != nil {
		return v1.NewRequestError(errors.New("get error "+err.Error()), http.StatusNotFound)
	}

	if post.IsEmpty() {
		return v1.NewRequestError(errors.New("get error"), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, post, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Get a post by ID
// @Description Get a specific post by its ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} post.Post
// @Failure 400 {object} v1.ErrorResponse
// @Failure 404 {object} v1.ErrorResponse
// @Router /posts/id/{id} [get]
func (h *Handler) GetById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	id, success := new(big.Int).SetString(params["id"], 10)
	if !success {
		return v1.NewRequestError(errors.New("id parse error"), http.StatusInternalServerError)
	}

	post, err := h.post.GetPostByID(ctx, id)
	if err != nil {
		return v1.NewRequestError(errors.New("get error "+err.Error()), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, post, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// @Summary Get all posts
// @Description Get all posts with pagination
// @Tags posts
// @Produce json
// @Param page query int false "Page number" default(0)
// @Param pageSize query int false "Page size" default(100)
// @Success 200 {array} post.Post
// @Failure 400 {object} v1.ErrorResponse
// @Failure 404 {object} v1.ErrorResponse
// @Router /posts/all [get]
func (h *Handler) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page, pageSize, err := extractPaginationParams(r)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	posts, err := h.post.Query(ctx, page, pageSize)
	if err != nil {
		return v1.NewRequestError(errors.New("get error "+err.Error()), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, posts, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

// extractPaginationParams extracts pagination parameters from the request query parameters
func extractPaginationParams(r *http.Request) (uint64, uint64, error) {
	const defaultPage uint64 = 0
	const defaultPageSize uint64 = 100

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil || pageStr == "" {
		page = defaultPage
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 64)
	if err != nil || pageSizeStr == "" {
		pageSize = defaultPageSize
	}

	return page, pageSize, nil
}

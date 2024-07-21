package postgrp

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strconv"

	"github.com/St3plox/Blogchain/business/core/post"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
)

type Handler struct {
	post *post.Core
	auth *auth.Auth
	user *user.Core
}

func New(postCore *post.Core, auth *auth.Auth, userCore *user.Core) *Handler {
	return &Handler{
		post: postCore,
		auth: auth,
		user: userCore,
	}
}

func (h *Handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var np post.NewPost
	err := json.NewDecoder(r.Body).Decode(&np)
	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusInternalServerError)
	}
	claims := auth.GetClaims(ctx)

	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusInternalServerError)
	}

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

func (h *Handler) GetUserPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims := auth.GetClaims(ctx)

	page, pageSize, err := extractPaginationParams(r)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusInternalServerError)
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

func (h *Handler) GetPostsByUserAddress(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	address := r.URL.Query().Get("address")

	page, pageSize, err := extractPaginationParams(r)
	if err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	posts, err := h.post.QueryByAddress(ctx, address, page, pageSize)
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

func (h *Handler) GetPostsByUserAddressAndIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	address := r.PathValue("address")
	index, err := strconv.ParseUint(r.PathValue("index"), 10, 64)
	if err != nil {
		return v1.NewRequestError(errors.New("parse error "+err.Error()), http.StatusInternalServerError)
	}

	post, err := h.post.QueryByIndex(ctx, address, index)
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

func (h *Handler) GetById(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	idStr := r.PathValue("id")
	id, success := new(big.Int).SetString(idStr, 10)
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

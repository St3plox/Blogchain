package postgrp

import (
	"context"
	"encoding/json"
	"errors"

	"net/http"

	"github.com/St3plox/Blogchain/business/core/post"
	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/google/uuid"
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

	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusInternalServerError)
	}

	usr, err := h.user.QueryByID(ctx, id)
	if err != nil {
		return v1.NewRequestError(errors.New("user error "+err.Error()), http.StatusNotFound)
	}

	post, err := h.post.Create(ctx, np, usr.AddressHex)
	if err != nil {
		return v1.NewRequestError(errors.New("Create error "+err.Error()), http.StatusInternalServerError)
	}

	err = web.Respond(ctx, w, post, http.StatusCreated)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) GetUserPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	claims := auth.GetClaims(ctx)

	id, err := uuid.Parse(claims.Subject)
	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusInternalServerError)
	}

	usr, err := h.user.QueryByID(ctx, id)
	if err != nil {
		return v1.NewRequestError(errors.New("user error "+err.Error()), http.StatusNotFound)
	}

	posts, err := h.post.QueryByAddress(ctx, usr.AddressHex)
	if err != nil {
		v1.NewRequestError(errors.New("post error "+err.Error()), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, posts, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) GetByUserAddressPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	
	address  := r.PathValue("address")

	posts, err := h.post.QueryByAddress(ctx, address)
	if err != nil {
		v1.NewRequestError(errors.New("post error "+err.Error()), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, posts, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) GetAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	posts, err := h.post.GetAllPostsSorted(ctx)
	if err != nil {
		v1.NewRequestError(errors.New("post error "+err.Error()), http.StatusNotFound)
	}

	err = web.Respond(ctx, w, posts, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

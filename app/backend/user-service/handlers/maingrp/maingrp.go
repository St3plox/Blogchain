package maingrp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/St3plox/Blogchain/business/core/user"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
)

type Handler struct {
	user *user.Core
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := web.Respond(ctx, w, nil, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) PostUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var nu user.NewUser
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		return v1.NewRequestError(errors.New("Decode error "+err.Error()), http.StatusInternalServerError)
	}

	usr, err := h.user.Create(ctx, nu)
	if err != nil {
		return v1.NewRequestError(errors.New("Create error "+err.Error()), http.StatusInternalServerError)
	}

	err = web.Respond(ctx, w, usr, http.StatusCreated)
	if err != nil {
		return err
	}

	return nil

}

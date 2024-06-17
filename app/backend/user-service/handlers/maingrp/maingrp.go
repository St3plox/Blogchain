package maingrp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
)

type Handler struct {
	user *user.Core
	auth *auth.Auth
}

func New(user *user.Core, auth *auth.Auth) *Handler {
	return &Handler{user: user, auth: auth}
}

func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := web.Respond(ctx, w, nil, http.StatusOK)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) RegisterUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var nu user.NewUser
	err := json.NewDecoder(r.Body).Decode(&nu)
	if err != nil {
		return v1.NewRequestError(errors.New("Decode error "+err.Error()), http.StatusInternalServerError)
	}

	usr, err := h.user.Create(ctx, nu)
	if err != nil {
		h.user.Delete(ctx, usr)
		return v1.NewRequestError(errors.New("Create error "+err.Error()), http.StatusInternalServerError)
	}

	roles := []user.Role{user.RoleUser}
	claims := auth.Claims{Roles: roles}
	token, err := h.auth.GenerateToken("private_key", claims)
	if err != nil {
		return v1.NewRequestError(errors.New("Token generation error "+err.Error()), http.StatusInternalServerError)
	}

	// Set JWT token in response header
	w.Header().Set("Authorization", "Bearer "+token)

	err = web.Respond(ctx, w, usr.Name, http.StatusCreated)
	if err != nil {
		h.user.Delete(ctx, usr)
		return err
	}

	return nil
}

func (h *Handler) LoginUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var credentials struct {
		Email          string `json:"email"`
		HashedPassword []byte `json:"hashed_password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		return v1.NewRequestError(errors.New("Decode error "+err.Error()), http.StatusBadRequest)
	}

	// Parse email address
	emailAddr, err := mail.ParseAddress(credentials.Email)
	if err != nil {
		return v1.NewRequestError(errors.New("invalid email address"), http.StatusBadRequest)
	}

	usr, err := h.user.QueryByEmail(ctx, *emailAddr)
	if err != nil {
		return v1.NewRequestError(errors.New("user not found"), http.StatusNotFound)
	}

	// Verify password
	if !bytes.Equal(usr.HashedPassword, credentials.HashedPassword) {
		return v1.NewRequestError(errors.New("invalid password"), http.StatusUnauthorized)
	}

	roles := []user.Role{user.RoleUser}

	// Generate JWT token
	claims := auth.Claims{Roles: roles}
	token, err := h.auth.GenerateToken(usr.ID.String(), claims)
	if err != nil {
		return v1.NewRequestError(errors.New("Token generation error "+err.Error()), http.StatusInternalServerError)
	}

	// Set JWT token in response header
	w.Header().Set("Authorization", "Bearer "+token)

	// Respond with user information (excluding sensitive data like password hash)
	respondUser := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}{
		ID:    usr.ID.String(),
		Email: usr.Email,
	}

	err = web.Respond(ctx, w, respondUser, http.StatusOK)
	if err != nil {
		return err
	}

	return nil
}

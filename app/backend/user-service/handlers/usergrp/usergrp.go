package usergrp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/mail"
	"time"

	"github.com/St3plox/Blogchain/business/core/user"
	"github.com/St3plox/Blogchain/business/web/auth"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	user *user.Core
	auth *auth.Auth
}

func New(user *user.Core, auth *auth.Auth) *Handler {
	return &Handler{user: user, auth: auth}
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
		return v1.NewRequestError(errors.New("Create error "+err.Error()), http.StatusBadRequest)
	}

	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Subject:   usr.ID.Hex(),
		},
		Roles: usr.Roles,
	}

	token, err := h.auth.GenerateToken("private_key", claims)
	if err != nil {
		return v1.NewRequestError(errors.New("Token generation error "+err.Error()), http.StatusInternalServerError)
	}

	// Set JWT token in response header
	w.Header().Set("Authorization", "Bearer "+token)

	err = web.Respond(ctx, w, user.Map(usr), http.StatusCreated)
	if err != nil {
		h.user.Delete(ctx, usr)
		return err
	}

	return nil
}

func (h *Handler) LoginUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var credentials user.UserCredentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		return v1.NewRequestError(errors.New("decode error "+err.Error()), http.StatusBadRequest)
	}

	// Parse email address
	emailAddr, err := mail.ParseAddress(credentials.Email)
	if err != nil {
		return v1.NewRequestError(errors.New("invalid email address"), http.StatusBadRequest)
	}

	h.user.Authenticate(ctx, *emailAddr, credentials.Password)

	usr, err := h.user.QueryByEmail(ctx, *emailAddr)
	if err != nil {
		return v1.NewRequestError(errors.New("user not found"), http.StatusNotFound)
	}

	// Verify password
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   usr.ID.Hex(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		Roles: usr.Roles,
	}

	token, err := h.auth.GenerateToken("private_key", claims)
	if err != nil {
		return v1.NewRequestError(errors.New("Token generation error "+err.Error()), http.StatusInternalServerError)
	}

	// Set JWT token in response header
	w.Header().Set("Authorization", "Bearer "+token)

	err = web.Respond(ctx, w, user.Map(usr), http.StatusOK)
	if err != nil {
		return err
	}

	return nil
}

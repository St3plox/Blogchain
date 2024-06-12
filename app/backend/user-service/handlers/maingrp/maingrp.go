package maingrp

import (
	"context"
	"net/http"

	"github.com/St3plox/Blogchain/foundation/web"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}


func (h *Handler) Get (ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	err := web.Respond(ctx, w, nil, http.StatusOK)
	if err != nil { 
		return err
	}
	return nil
}
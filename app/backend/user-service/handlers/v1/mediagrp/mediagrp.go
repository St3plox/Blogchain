package mediagrp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/St3plox/Blogchain/business/core/media"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gorilla/mux"
)

type Handler struct {
	media *media.Core
}

func New(media *media.Core) *Handler {
	return &Handler{
		media: media,
	}
}

func (h *Handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Parse the multipart form, allowing for a maximum upload MaxFileSizeMb
	err := r.ParseMultipartForm(h.media.MaxFileSizeMb << 20)
	if err != nil {
		return v1.NewRequestError(errors.New("failed to parse multipart form: "+err.Error()), http.StatusBadRequest)
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return v1.NewRequestError(errors.New("read file error: "+err.Error()), http.StatusBadRequest)
	}
	defer file.Close()

	buf := make([]byte, handler.Size)
	_, err = file.Read(buf)
	if err != nil {
		return v1.NewRequestError(errors.New("error reading file: "+err.Error()), http.StatusInternalServerError)
	}

	// Check if the file is an image
	mime := mimetype.Detect(buf)
	if !mime.Is("image/jpeg") && !mime.Is("image/png") && !mime.Is("image/gif") && !mime.Is("image/bmp") {
		return v1.NewRequestError(errors.New("file is not a valid image type"), http.StatusBadRequest)
	}

	newMedia := media.NewMedia{
		Filename:  handler.Filename,
		Length:    handler.Size,
		FileBytes: buf,
	}

	media, err := h.media.Create(ctx, newMedia)
	if err != nil {
		return v1.NewRequestError(errors.New("error posting media: "+err.Error()), http.StatusInternalServerError)
	}

	err = web.Respond(ctx, w, media, http.StatusCreated)
	if err != nil {
		h.media.Delete(ctx, media)
		return err
	}

	return nil
}

func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	params := mux.Vars(r)

	mediaVal, err := h.media.QueryByID(ctx, params[MediaID])
	if err != nil {
		if errors.Is(err, media.ErrNotFound) {

			return v1.NewRequestError(errors.New("media not found"), http.StatusNotFound)
		}
		return v1.NewRequestError(errors.New("error getting media: "+err.Error()), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, mediaVal, http.StatusOK)
}

func (h *Handler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	params := mux.Vars(r)

	err := h.media.DeleteByID(ctx, params[MediaID])
	if err != nil {
		if errors.Is(err, media.ErrNotFound) {
			return v1.NewRequestError(errors.New("media not found"), http.StatusNotFound)
		}
		return v1.NewRequestError(errors.New("error deleting media: "+err.Error()), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// GetByIDs handles the retrieval of multiple media files by their IDs.
func (h *Handler) GetByIDs(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var mediaIDs []string
	if err := json.NewDecoder(r.Body).Decode(&mediaIDs); err != nil {
		return v1.NewRequestError(errors.New("invalid request payload: "+err.Error()), http.StatusBadRequest)
	}

	medias, err := h.media.QueryByIDs(ctx, mediaIDs)
	if err != nil {
		return v1.NewRequestError(errors.New("error getting medias: "+err.Error()), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, medias, http.StatusOK)
}

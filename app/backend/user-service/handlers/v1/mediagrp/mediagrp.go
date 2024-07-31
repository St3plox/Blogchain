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

// Post handles the uploading of a new media file.
// @Summary Upload a media file
// @Description Upload a new media file (image only)
// @Tags media
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "File to upload"
// @Success 201 {object} media.Media
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /media [post]
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

// Get handles the retrieval of a media file by its ID.
// @Summary Get a media file
// @Description Retrieve a media file by its ID
// @Tags media
// @Produce application/json
// @Param media_id path string true "Media ID"
// @Success 200 {object} media.Media
// @Failure 404 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /media/{media_id} [get]
func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	mediaVal, err := h.media.QueryByID(ctx, params["media_id"])
	if err != nil {
		if errors.Is(err, media.ErrNotFound) {
			return v1.NewRequestError(errors.New("media not found"), http.StatusNotFound)
		}
		return v1.NewRequestError(errors.New("error getting media: "+err.Error()), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, mediaVal, http.StatusOK)
}

// Delete handles the deletion of a media file by its ID.
// @Summary Delete a media file
// @Description Delete a media file by its ID
// @Tags media
// @Param media_id path string true "Media ID"
// @Success 204
// @Failure 404 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /media/{media_id} [delete]
func (h *Handler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	err := h.media.DeleteByID(ctx, params["media_id"])
	if err != nil {
		if errors.Is(err, media.ErrNotFound) {
			return v1.NewRequestError(errors.New("media not found"), http.StatusNotFound)
		}
		return v1.NewRequestError(errors.New("error deleting media: "+err.Error()), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// GetByIDs handles the retrieval of multiple media files by their IDs.
// @Summary Get multiple media files
// @Description Retrieve multiple media files by their IDs
// @Tags media
// @Accept application/json
// @Produce application/json
// @Param ids body []string true "Array of Media IDs"
// @Success 200 {array} media.Media
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /media [get]
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

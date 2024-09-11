package mediagrp

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/St3plox/Blogchain/business/core/media"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
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
// @Success 201 {object} media.MediaData
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/media [post]
func (h *Handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	newMedia, err := h.media.ParseMedia(r)
	if err != nil {
		return err
	}

	media, err := h.media.Create(ctx, newMedia)
	if err != nil {
		return v1.NewRequestError(errors.New("error posting media: "+err.Error()), http.StatusInternalServerError)
	}

	err = web.Respond(ctx, w, media, http.StatusCreated)
	if err != nil {
		return err
	}

	return nil
}

// Post handles the uploading of multiple new media files.
// @Summary Upload multiple media file
// @Description Upload a new media file (image only)
// @Tags media
// @Accept multipart/form-data
// @Produce application/json
// @Param file formData file true "File to upload"
// @Success 201 {object} []media.MediaData	
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/media/many [post]
func (h *Handler) PostMultiple(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	newMedia, err := h.media.ParseMultipleMedia(r)
	if err != nil {
		return err
	}

	media, err := h.media.CreateMultiple(ctx, newMedia)
	if err != nil {
		return v1.NewRequestError(errors.New("error posting media: "+err.Error()), http.StatusInternalServerError)
	}

	err = web.Respond(ctx, w, media, http.StatusCreated)
	if err != nil {
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
// @Router /v1/media/{media_id} [get]
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

// Delete handles the deletion of a media file by its ID.
// @Summary Delete a media file
// @Description Delete a media file by its ID
// @Tags media
// @Param media_id path string true "Media ID"
// @Success 204
// @Failure 404 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/media/{media_id} [delete]
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
// @Summary Get multiple media files
// @Description Retrieve multiple media files by their IDs
// @Tags media
// @Accept application/json
// @Produce application/json
// @Param ids body []string true "Array of Media IDs"
// @Success 200 {array} media.Media
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/media [get]
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

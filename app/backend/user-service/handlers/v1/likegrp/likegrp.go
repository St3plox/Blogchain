package likegrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/St3plox/Blogchain/business/core/like"
	"github.com/St3plox/Blogchain/business/core/like/likedb"
	v1 "github.com/St3plox/Blogchain/business/web/v1"
	"github.com/St3plox/Blogchain/foundation/web"
	"github.com/gorilla/mux"
)

// Handler manages the set of like endpoints.
type Handler struct {
	like *like.Core
}

// New creates a handler for like-related requests.
func New(likeCore *like.Core) *Handler {
	return &Handler{
		like: likeCore,
	}
}

// Get handles the retrieval of a like by its ID.
// @Summary Get a like
// @Description Retrieve a like by its ID
// @Tags likes
// @Produce application/json
// @Param id path string true "Like ID"
// @Success 200 {object} like.Like
// @Failure 404 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/like/{id} [get]
func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		return v1.NewRequestError(errors.New("no id param in request"), http.StatusBadRequest)
	}

	foundLike, err := h.like.QueryByID(ctx, id)
	if err != nil {
		if likedb.IsLikeNotFound(err) {
			return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusNotFound)
		}
		return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, foundLike, http.StatusOK)
}

// GetByUserOrPostID handles the retrieval of likes by user ID or post ID.
// @Summary Get likes by user ID or post ID
// @Description Retrieve likes by either user ID or post ID
// @Tags likes
// @Produce application/json
// @Param user_id query string false "User ID"
// @Param post_id query string false "Post ID"
// @Success 200 {array} like.Like
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/like [get]
func (h *Handler) GetByUserOrPostID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	userID := query.Get("user_id")
	postID := query.Get("post_id")

	var likes []like.Like
	var err error

	if userID != "" {
		likes, err = h.like.QueryAllByUserID(ctx, userID)
	} else if postID != "" {
		likes, err = h.like.QueryAllByPostID(ctx, postID)
	} else {
		return v1.NewRequestError(errors.New("user_id or post_id must be provided"), http.StatusBadRequest)
	}

	if err != nil {
		return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, likes, http.StatusOK)
}

// Create handles creating a new like.
// @Summary Create a like
// @Description Create a new like
// @Tags likes
// @Accept application/json
// @Produce application/json
// @Param like body like.Like true "Like object"
// @Success 201 {object} like.Like
// @Failure 400 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/like [post]
func (h *Handler) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var newLike like.Like
	if err := web.Decode(r, &newLike); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}

	createdLike, err := h.like.Create(ctx, newLike)
	if err != nil {
		return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, createdLike, http.StatusCreated)
}

// Update handles updating an existing like.
// @Summary Update a like
// @Description Update an existing like by its ID
// @Tags likes
// @Accept application/json
// @Produce application/json
// @Param id path string true "Like ID"
// @Param like body like.Like true "Updated like object"
// @Success 200 {object} like.Like
// @Failure 400 {object} v1.ErrorResponse
// @Failure 404 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/like/{id} [put]
func (h *Handler) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var updatedLike like.Like
	if err := web.Decode(r, &updatedLike); err != nil {
		return v1.NewRequestError(err, http.StatusBadRequest)
	}
	
	like, err := h.like.Update(ctx, updatedLike)
	if err != nil {
		if likedb.IsLikeNotFound(err) {
			return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusNotFound)
		}
		return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, like, http.StatusOK)
}

// Delete handles deleting an existing like by ID.
// @Summary Delete a like
// @Description Delete a like by its ID
// @Tags likes
// @Produce application/json
// @Param id path string true "Like ID"
// @Success 204 "No Content"
// @Failure 400 {object} v1.ErrorResponse
// @Failure 404 {object} v1.ErrorResponse
// @Failure 500 {object} v1.ErrorResponse
// @Router /v1/like/{id} [delete]
func (h *Handler) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		return v1.NewRequestError(errors.New("no id param in request"), http.StatusBadRequest)
	}

	if err := h.like.DeleteByID(ctx, id); err != nil {
		if likedb.IsLikeNotFound(err) {
			return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusNotFound)
		}
		return v1.NewRequestError(fmt.Errorf("core error: %w", err), http.StatusInternalServerError)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

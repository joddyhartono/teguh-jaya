package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/joddyhartono/teguh-jaya/internal/model"
	"github.com/joddyhartono/teguh-jaya/internal/response"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]model.Category, error)
	Delete(ctx context.Context, id uuid.UUID) (error)
}

type CategoryHandler struct {
	repo CategoryRepository
}

func NewCategoryHandler(repo CategoryRepository) (*CategoryHandler) {
	return &CategoryHandler{
		repo: repo,
	}
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.List(r.Context())
	if err != nil {
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	res := make([]categoryListResponse, 0)
	for _, category := range categories {
		res = append(res, toCategoryListResponse(category))
	}
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request ) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "invalid uuid")
		return
	}

	err = h.repo.Delete(r.Context(), id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			response.WriteJSONError(w, http.StatusNotFound, "category not found")
			return
		}
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
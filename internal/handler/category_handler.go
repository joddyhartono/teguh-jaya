package handler

import (
	"context"
	"net/http"

	"github.com/joddyhartono/teguh-jaya/internal/model"
	"github.com/joddyhartono/teguh-jaya/internal/response"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]model.Category, error)
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
	response.WriteJSONData(w, http.StatusOK, res)
}
package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/joddyhartono/teguh-jaya/internal/model"
	"github.com/joddyhartono/teguh-jaya/internal/response"
)

type ProductRepository interface {
	List(ctx context.Context) ([]model.Product, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error)
}

type ProductHandler struct {
	repo ProductRepository
}

func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{
		repo: repo,
	}
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.List(r.Context())
	if err != nil {
		log.Printf("failed to list products: %v", err)
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	res := make([]productListResponse, 0)
	for _, product := range products {
		res = append(res, toProductListResponse(product))
	}
	response.WriteJSON(w, http.StatusOK, res)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		response.WriteJSONError(w, http.StatusBadRequest, "invalid id")
		return
	}
	product, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		log.Printf("failed to get product: %v", err)
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	response.WriteJSON(w, http.StatusOK, toProductDetailResponse(*product))
}
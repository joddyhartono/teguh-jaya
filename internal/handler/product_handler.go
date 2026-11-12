package handler

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/joddyhartono/teguh-jaya/internal/model"
	"github.com/joddyhartono/teguh-jaya/internal/response"
)

type ProductService interface {
	List(ctx context.Context) ([]model.Product, error)
}

type ProductHandler struct {
	svc ProductService
}

func NewProductHandler(svc ProductService) *ProductHandler {
	h := ProductHandler{
		svc: svc,
	}
	return &h
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	slog.Info("List dipanggil")
	products, err := h.svc.List(r.Context())
	if err != nil {
		slog.Error("failed to list products", "error", err)
		response.WriteJSONError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	res := make([]productResponse, 0)
	for _, p := range products {
		res = append(res, toProductListResponse(p))
	}
	response.WriteJSONData(w, http.StatusOK, res)
}
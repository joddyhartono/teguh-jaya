package handler

import (
	"github.com/google/uuid"
	"github.com/joddyhartono/teguh-jaya/internal/model"
)

type productListResponse struct {
	ID			uuid.UUID	`json:"id"`
	CategoryID	uuid.UUID	`json:"category_id"`
	Name		string		`json:"name"`
	Image		string		`json:"image"`
}

type productDetailResponse struct {
	ID			uuid.UUID	`json:"id"`
	CategoryID	uuid.UUID	`json:"category_id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
}

func toProductListResponse(product model.Product) productListResponse {
	return productListResponse{
		ID: product.ID,
		CategoryID: product.CategoryID,
		Name: product.Name,
		Image: product.Image,
	}
}

func toProductDetailResponse(product model.Product) productDetailResponse {
	return productDetailResponse{
		ID: product.ID,
		CategoryID: product.CategoryID,
		Name: product.Name,
		Description: product.Description,
		Image: product.Image,
	}
}
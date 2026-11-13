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

func toProductListResponse(p model.Product) productListResponse {
	return productListResponse{
		ID: p.ID,
		CategoryID: p.CategoryID,
		Name: p.Name,
		Image: p.Image,
	}
}

func toProductDetailResponse(p model.Product) productDetailResponse {
	return productDetailResponse{
		ID: p.ID,
		CategoryID: p.CategoryID,
		Name: p.Name,
		Description: p.Description,
		Image: p.Image,
	}
}
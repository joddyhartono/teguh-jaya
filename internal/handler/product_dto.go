package handler

import (
	"github.com/google/uuid"
	"github.com/joddyhartono/teguh-jaya/internal/model"
)

type productResponse struct {
	ID			uuid.UUID	`json:"id"`
	CategoryID	uuid.UUID	`json:"category_id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
}

func toProductListResponse(p model.Product) productResponse {
	return productResponse{
		ID: p.ID,
		CategoryID: p.CategoryID,
		Name: p.Name,
		Description: p.Description,
		Image: p.Image,
	}
}
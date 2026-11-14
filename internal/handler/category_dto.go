package handler

import (
	"github.com/google/uuid"
	"github.com/joddyhartono/teguh-jaya/internal/model"
)

type categoryListResponse struct {
	ID			uuid.UUID	`json:"id"`
	Name		string		`json:"name"`
	Description	string		`json:"description"`
	Image		string		`json:"image"`
}

func toCategoryListResponse(category model.Category) categoryListResponse {
	return categoryListResponse{
		ID: category.ID,
		Name: category.Name,
		Description: category.Description,
		Image: category.Image,
	}
}
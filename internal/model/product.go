package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID			uuid.UUID
	CategoryID	uuid.UUID
	Name		string
	Description	string
	Image		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	RowColumn	int16
}
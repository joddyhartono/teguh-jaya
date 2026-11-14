package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID			uuid.UUID
	Name		string
	Description	string
	Image		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	RowColumn	int16
}
package repository

import (
	"context"
	"database/sql"

	"github.com/joddyhartono/teguh-jaya/internal/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) List(ctx context.Context) ([]model.Category, error) {
	query := `
		SELECT
			id,
			name,
			description,
			image
		FROM categories
		WHERE row_column = 0
	`

	categories := make([]model.Category, 0)
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category model.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.Image)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, rows.Err()
}
package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
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

func (r *CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE categories
		SET row_column = 1
		WHERE row_column = 0 AND id = $1
	`

	res, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return model.ErrNotFound
	}

	return nil
}
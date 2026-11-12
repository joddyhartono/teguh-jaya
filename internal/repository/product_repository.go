package repository

import (
	"context"
	"database/sql"

	"github.com/joddyhartono/teguh-jaya/internal/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) List(ctx context.Context) ([]model.Product, error) {
	query := `
		SELECT
			id,
			category_id,
			name,
			description,
			image
		FROM products
		WHERE row_column = 0
	`

	products := make([]model.Product, 0)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		err = rows.Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.Image)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, rows.Err()
}
package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
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
			category_id
			name,
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
		var product model.Product
		err = rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Description, &product.Image)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, rows.Err()
}


func (r *ProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Product, error) {
	query := `
		SELECT
			id,
			category_id,
			name,
			description,
			image
		FROM products
		WHERE id = $1 AND row_column = 0
	`

	var product model.Product
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.CategoryID,
		&product.Name,
		&product.Description,
		&product.Image,
	)
	
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}
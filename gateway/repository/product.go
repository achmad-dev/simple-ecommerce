package repository

import (
	"database/sql"
	"errors"

	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
)

type ProductRepository interface {
	FetchAllProducts() ([]domain.Product, error)
	FetchProductsPaginated(limit, offset int) ([]domain.Product, error)
	FetchProductByName(name string) ([]domain.Product, error)
	GetProductByID(id int) (*domain.Product, error)
}

type baseProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &baseProductRepository{db: db}
}

func (r *baseProductRepository) FetchAllProducts() ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, title, price, description FROM simple_ecommerce.product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.Price, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *baseProductRepository) FetchProductsPaginated(limit, offset int) ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, title, price, description FROM simple_ecommerce.product LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.Price, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *baseProductRepository) FetchProductByName(name string) ([]domain.Product, error) {
	rows, err := r.db.Query("SELECT id, title, price, description FROM simple_ecommerce.product WHERE title ILIKE '%' || $1 || '%'", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Title, &product.Price, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *baseProductRepository) GetProductByID(id int) (*domain.Product, error) {
	var product domain.Product
	err := r.db.QueryRow("SELECT id, title, price, description FROM simple_ecommerce.product WHERE id = $1", id).Scan(&product.ID, &product.Title, &product.Price, &product.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found") // No product found with the given ID
		}
		return nil, err
	}
	return &product, nil
}

package repository

import (
	"database/sql"

	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
)

type CartRepository interface {
	AddProductToCart(userID, productID int) error
	RemoveProductFromCart(userID, productID int) error
	GetCartByUserID(userID int) ([]domain.CartItem, error)
}

type baseCartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &baseCartRepository{db: db}
}

func (r *baseCartRepository) AddProductToCart(userID, productID int) error {
	_, err := r.db.Exec("INSERT INTO simple_ecommerce.cart (user_id, product_id) VALUES ($1, $2)", userID, productID)
	return err
}

func (r *baseCartRepository) RemoveProductFromCart(userID, productID int) error {
	_, err := r.db.Exec("DELETE FROM simple_ecommerce.cart WHERE user_id = $1 AND product_id = $2", userID, productID)
	return err
}

func (r *baseCartRepository) GetCartByUserID(userID int) ([]domain.CartItem, error) {
	rows, err := r.db.Query("SELECT id, product_id FROM simple_ecommerce.cart WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []domain.CartItem
	for rows.Next() {
		var cartItem domain.CartItem
		if err := rows.Scan(&cartItem.Id, &cartItem.ProductID); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

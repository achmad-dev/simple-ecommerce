package repository

import (
	"database/sql"

	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/lib/pq"
)

type OrderRepository interface {
	CreateOrder(userID int, cartIDs []int, paymentMethodID int, totalPrice int) error
	GetOrdersByUserID(userID int) ([]domain.Order, error)
	UpdateOrderToPaid(orderID int) error
}

type baseOrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) OrderRepository {
	return &baseOrderRepository{db: db}
}

func (r *baseOrderRepository) CreateOrder(userID int, cartIDs []int, paymentMethodID int, totalPrice int) error {
	_, err := r.db.Exec("INSERT INTO simple_ecommerce.order (user_id, cart_ids, payment_method_id, total_price) VALUES ($1, $2, $3, $4)", userID, pq.Array(cartIDs), paymentMethodID, totalPrice)
	return err
}

func (r *baseOrderRepository) GetOrdersByUserID(userID int) ([]domain.Order, error) {
	rows, err := r.db.Query("SELECT id, user_id, cart_ids, payment_method_id, total_price, is_paid, created_at, updated_at FROM simple_ecommerce.order WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order.ID, &order.UserID, pq.Array(&order.CartIDs), &order.PaymentMethodID, &order.TotalPrice, &order.IsPaid, &order.CreatedAt, &order.UpdatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *baseOrderRepository) UpdateOrderToPaid(orderID int) error {
	_, err := r.db.Exec("UPDATE simple_ecommerce.order SET is_paid = TRUE WHERE id = $1", orderID)
	return err
}

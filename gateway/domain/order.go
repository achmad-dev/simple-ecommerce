package domain

import "time"

type Order struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	CartIDs         []int     `json:"cart_ids"`
	PaymentMethodID int       `json:"payment_method_id"`
	TotalPrice      int       `json:"total_price"`
	IsPaid          bool      `json:"is_paid"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

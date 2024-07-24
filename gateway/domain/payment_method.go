package domain

type PaymentMethod struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	BankName string `json:"bank_name"`
}

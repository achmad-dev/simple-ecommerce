package dto

type PaymentMethodDto struct {
	UserID   int    `json:"user_id,omitempty"`
	BankName string `json:"bank_name"`
}

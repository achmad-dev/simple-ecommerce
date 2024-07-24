package repository

import (
	"database/sql"

	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/dto"
)

type PaymentRepository interface {
	CreatePaymentMethod(paymentDto dto.PaymentMethodDto) error
	GetPaymentMethodsByUserID(userID int) ([]domain.PaymentMethod, error)
}

type basePaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &basePaymentRepository{db: db}
}

func (r *basePaymentRepository) CreatePaymentMethod(paymentDto dto.PaymentMethodDto) error {
	_, err := r.db.Exec(
		"INSERT INTO simple_ecommerce.payment_method (user_id, bank_name) VALUES ($1, $2)",
		paymentDto.UserID, paymentDto.BankName,
	)
	return err
}

func (r *basePaymentRepository) GetPaymentMethodsByUserID(userID int) ([]domain.PaymentMethod, error) {
	rows, err := r.db.Query("SELECT id, user_id, bank_name FROM simple_ecommerce.payment_method WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []domain.PaymentMethod
	for rows.Next() {
		var paymentMethod domain.PaymentMethod
		if err := rows.Scan(&paymentMethod.ID, &paymentMethod.UserID, &paymentMethod.BankName); err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}
	return paymentMethods, nil
}

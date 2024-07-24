package service

import (
	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/dto"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/sirupsen/logrus"
)

type PaymentService interface {
	CreatePaymentMethod(paymentDto dto.PaymentMethodDto) error
	GetPaymentMethodsByUserID(userID int) ([]domain.PaymentMethod, error)
}

type basePaymentService struct {
	paymentRepo repository.PaymentRepository
	log         *logrus.Logger
}

func NewPaymentService(paymentRepo repository.PaymentRepository, log *logrus.Logger) PaymentService {
	return &basePaymentService{paymentRepo: paymentRepo, log: log}
}

func (s *basePaymentService) CreatePaymentMethod(paymentDto dto.PaymentMethodDto) error {
	s.log.Infof("Creating payment method for user %d", paymentDto.UserID)

	err := s.paymentRepo.CreatePaymentMethod(paymentDto)
	if err != nil {
		s.log.Errorf("Error creating payment method for user %d: %v", paymentDto.UserID, err)
		return err
	}

	s.log.Infof("Payment method created successfully for user %d", paymentDto.UserID)
	return nil
}

func (s *basePaymentService) GetPaymentMethodsByUserID(userID int) ([]domain.PaymentMethod, error) {
	s.log.Infof("Getting payment methods for user %d", userID)

	paymentMethods, err := s.paymentRepo.GetPaymentMethodsByUserID(userID)
	if err != nil {
		s.log.Errorf("Error getting payment methods for user %d: %v", userID, err)
		return nil, err
	}

	s.log.Infof("Retrieved %d payment methods for user %d", len(paymentMethods), userID)
	return paymentMethods, nil
}

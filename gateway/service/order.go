package service

import (
	"errors"

	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/sirupsen/logrus"
)

type OrderService interface {
	CreateOrder(email string) error
	GetOrdersByUserEmail(email string) ([]domain.Order, error)
	PayOrder(orderID int) error
}

type baseOrderService struct {
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	userRepo    repository.UserRepository
	paymentRepo repository.PaymentRepository
	productRepo repository.ProductRepository
	log         *logrus.Logger
}

func NewOrderService(orderRepo repository.OrderRepository, cartRepo repository.CartRepository, userRepo repository.UserRepository, paymentRepo repository.PaymentRepository, productRepo repository.ProductRepository, log *logrus.Logger) OrderService {
	return &baseOrderService{orderRepo: orderRepo, cartRepo: cartRepo, userRepo: userRepo, paymentRepo: paymentRepo, productRepo: productRepo, log: log}
}

func (s *baseOrderService) CreateOrder(email string) error {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	// if user doesn't have payment method the it's not processed
	payment, err := s.paymentRepo.GetPaymentMethodsByUserID(user.Id)
	if err != nil {
		return err
	}
	//use default payment (for now the first in the array)
	s.log.Infof("Creating order for user %d with payment method %d", user.Id, payment[0].ID)

	var totalPrice int
	carts, err := s.cartRepo.GetCartByUserID(user.Id)
	if err != nil {
		s.log.Errorf("Error getting cart by user ID %d: %v", user.Id, err)
		return err
	}
	if len(carts) == 0 {
		return errors.New("add product to cart first")
	}

	cartIds := make([]int, len(carts))
	for i, cartItem := range carts {
		product, err := s.productRepo.GetProductByID(cartItem.ProductID)
		if err != nil {
			s.log.Errorf("Error getting product by ID %d: %v", cartItem.ProductID, err)
			return err
		}
		totalPrice += product.Price
		cartIds[i] = cartItem.Id
	}

	err = s.orderRepo.CreateOrder(user.Id, cartIds, payment[0].ID, totalPrice)
	if err != nil {
		s.log.Errorf("Error creating order for user %d: %v", user.Id, err)
		return err
	}

	s.log.Infof("Order created for user %d with total price %d", user.Id, totalPrice)
	return nil
}

func (s *baseOrderService) GetOrdersByUserEmail(email string) ([]domain.Order, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	s.log.Infof("Getting orders for user %d", user.Id)

	orders, err := s.orderRepo.GetOrdersByUserID(user.Id)
	if err != nil {
		s.log.Errorf("Error getting orders for user %d: %v", user.Id, err)
		return nil, err
	}

	s.log.Infof("Retrieved %d orders for user %d", len(orders), user.Id)
	return orders, nil
}

func (s *baseOrderService) PayOrder(orderID int) error {
	s.log.Infof("Paying order %d", orderID)

	err := s.orderRepo.UpdateOrderToPaid(orderID)
	if err != nil {
		s.log.Errorf("Error paying order %d: %v", orderID, err)
		return err
	}

	s.log.Infof("Order %d paid successfully", orderID)
	return nil
}

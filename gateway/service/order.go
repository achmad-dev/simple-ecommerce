package service

import (
	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/sirupsen/logrus"
)

type OrderService interface {
	CreateOrder(userID int, paymentMethodID int) error
	GetOrdersByUserID(userID int) ([]domain.Order, error)
	PayOrder(orderID int) error
}

type baseOrderService struct {
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
	log         *logrus.Logger
}

func NewOrderService(orderRepo repository.OrderRepository, cartRepo repository.CartRepository, productRepo repository.ProductRepository, log *logrus.Logger) OrderService {
	return &baseOrderService{orderRepo: orderRepo, cartRepo: cartRepo, productRepo: productRepo, log: log}
}

func (s *baseOrderService) CreateOrder(userID int, paymentMethodID int) error {
	s.log.Infof("Creating order for user %d with payment method %d", userID, paymentMethodID)

	var totalPrice int
	carts, err := s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		s.log.Errorf("Error getting cart by user ID %d: %v", userID, err)
		return err
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

	err = s.orderRepo.CreateOrder(userID, cartIds, paymentMethodID, totalPrice)
	if err != nil {
		s.log.Errorf("Error creating order for user %d: %v", userID, err)
		return err
	}

	s.log.Infof("Order created for user %d with total price %d", userID, totalPrice)
	return nil
}

func (s *baseOrderService) GetOrdersByUserID(userID int) ([]domain.Order, error) {
	s.log.Infof("Getting orders for user %d", userID)

	orders, err := s.orderRepo.GetOrdersByUserID(userID)
	if err != nil {
		s.log.Errorf("Error getting orders for user %d: %v", userID, err)
		return nil, err
	}

	s.log.Infof("Retrieved %d orders for user %d", len(orders), userID)
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

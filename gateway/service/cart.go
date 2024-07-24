package service

import (
	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/sirupsen/logrus"
)

type CartService interface {
	AddProductToCart(email string, productID int) error
	RemoveProductFromCart(email string, productID int) error
	GetCartByUserID(userID int) ([]domain.CartItem, error)
}

type baseCartService struct {
	cartRepo repository.CartRepository
	userRepo repository.UserRepository
	log      *logrus.Logger
}

func NewCartService(cartRepo repository.CartRepository, userRepo repository.UserRepository, log *logrus.Logger) CartService {
	return &baseCartService{cartRepo: cartRepo, userRepo: userRepo, log: log}
}

func (b *baseCartService) AddProductToCart(email string, productID int) error {
	user, err := b.userRepo.GetUserByEmail(email)
	if err != nil {
		b.log.Errorf("error adding product to cart: %v", err)
		return err
	}
	err = b.cartRepo.AddProductToCart(user.Id, productID)
	if err != nil {
		b.log.Errorf("error adding product to cart: %v", err)
		return err
	}
	b.log.Infof("product with ID %d added to cart for user %d", productID, user.Id)
	return nil
}

func (b *baseCartService) RemoveProductFromCart(email string, productID int) error {
	user, err := b.userRepo.GetUserByEmail(email)
	if err != nil {
		b.log.Errorf("error remove product to cart: %v", err)
		return err
	}
	err = b.cartRepo.RemoveProductFromCart(user.Id, productID)
	if err != nil {
		b.log.Errorf("error removing product from cart: %v", err)
		return err
	}
	b.log.Infof("product with ID %d removed from cart for user %d", productID, user.Id)
	return nil
}

func (b *baseCartService) GetCartByUserID(userID int) ([]domain.CartItem, error) {
	cartItems, err := b.cartRepo.GetCartByUserID(userID)
	if err != nil {
		b.log.Errorf("error getting cart items for user %d: %v", userID, err)
		return nil, err
	}
	b.log.Infof("retrieved cart items for user %d", userID)
	return cartItems, nil
}

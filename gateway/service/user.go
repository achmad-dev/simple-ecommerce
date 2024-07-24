package service

import (
	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/dto"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	RegisterUser(user dto.AuthUserDto) error
	GetUserByEmail(email string) (*domain.User, error)
}

type baseUserService struct {
	userRepo repository.UserRepository
	log      *logrus.Logger
}

func NewUserService(userRepo repository.UserRepository, log *logrus.Logger) UserService {
	return &baseUserService{userRepo: userRepo, log: log}
}

func (b *baseUserService) RegisterUser(user dto.AuthUserDto) error {
	b.log.Infof("Registering user with email: %s", user.Email)
	err := b.userRepo.RegisterUser(user)
	if err != nil {
		b.log.Errorf("Error registering user with email %s: %v", user.Email, err)
		return err
	}
	b.log.Infof("User registered successfully with email: %s", user.Email)
	return nil
}

func (b *baseUserService) GetUserByEmail(email string) (*domain.User, error) {
	b.log.Infof("Getting user by email: %s", email)
	user, err := b.userRepo.GetUserByEmail(email)
	if err != nil {
		b.log.Errorf("Error getting user by email %s: %v", email, err)
		return nil, err
	}
	b.log.Infof("User retrieved successfully with email: %s", email)
	return user, nil
}

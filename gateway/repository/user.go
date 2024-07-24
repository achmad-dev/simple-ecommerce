package repository

import (
	"database/sql"

	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	RegisterUser(user_dto dto.AuthUserDto) error
	GetUserByEmail(email string) (*domain.User, error)
}

type baseUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &baseUserRepository{db: db}
}

func (b *baseUserRepository) RegisterUser(user_dto dto.AuthUserDto) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user_dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = b.db.Exec("INSERT INTO simple_ecommerce.user (username, password) VALUES ($1, $2)", user_dto.Email, string(hashedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (b *baseUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var storedUser domain.User
	err := b.db.QueryRow("SELECT id, username, password FROM simple_ecommerce.user WHERE username = $1", email).Scan(&storedUser.Id, &storedUser.Email, &storedUser.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &storedUser, nil
}

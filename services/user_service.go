package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ropc-service/model/entities"
	"ropc-service/repositories"
)

type UserService interface {
	CreateUser(user *entities.User) (*entities.User, error)
}

type DefaultUserService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *DefaultUserService {
	return &DefaultUserService{repository: repository}
}

func (selfC DefaultUserService) CreateUser(user *entities.User) (*entities.User, error) {
	if user.Username == "" {
		return nil, errors.New("username is required")
	}

	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	user.Password = string(hashedPassword)

	newUser, err := selfC.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

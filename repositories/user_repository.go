package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ropc-service/conf"
	"ropc-service/model/entities"
)

type UserRepository interface {
	GetUser(username string) (*entities.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: conf.DB,
	}
}

func (selfC UserRepositoryImpl) GetUser(username string) (*entities.User, error) {
	var user entities.User

	err := selfC.db.Model(&entities.User{}).Where("username = ?", username).Preload("Client").First(&user).Error
	if err != nil {
		return nil, errors.New("invalid user credentials")
	}

	return &user, nil
}

func (selfC UserRepositoryImpl) GetUserAndClient(u *entities.User, client *entities.Client) (*entities.User, error) {

	var user *entities.User

	err := selfC.db.Where("username = ? and client_id = ?", u.Username, client.ClientId).First(&user).Error

	if err != nil {
		return nil, errors.New("invalid credential")
	}

	return user, nil
}

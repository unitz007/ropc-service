package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ropc-service/conf"
	"ropc-service/model"
)

type UserRepository interface {
	GetUser(username string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: conf.DB,
	}
}

func (selfC UserRepositoryImpl) GetUser(username string) (*model.User, error) {
	var user model.User

	err := selfC.db.Model(&model.User{}).Where("username = ?", username).Preload("Client").First(&user).Error
	if err != nil {
		return nil, errors.New("invalid user credentials")
	}

	return &user, nil
}

func (selfC UserRepositoryImpl) GetUserAndClient(u *model.User, client *model.Client) (*model.User, error) {

	var user *model.User

	err := selfC.db.Where("username = ? and client_id = ?", u.Username, client.ClientId).First(&user).Error

	if err != nil {
		return nil, errors.New("invalid credential")
	}

	return user, nil
}

package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ropc-service/conf"
	"ropc-service/model/entities"
)

type UserRepository interface {
	GetUser(username string) (*entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
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

	err := selfC.db.Model(&entities.User{}).Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, errors.New("invalid user credentials")
	}

	return &user, nil
}

func (selfC UserRepositoryImpl) CreateUser(user *entities.User) (*entities.User, error) {

	err := selfC.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"ropc-service/conf"
	"ropc-service/model/entities"
)

type UserRepository interface {
	GetUser(usernameOrEmail string) (*entities.User, error)
	GetUserByUsernameOrEmail(username, email string) (*entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
}

type userRepository struct {
	db conf.Database[gorm.DB]
}

func NewUserRepository(database conf.Database[gorm.DB]) UserRepository {
	return &userRepository{
		db: database,
	}
}

func (selfC userRepository) GetUser(username string) (*entities.User, error) {
	var user *entities.User

	err := selfC.db.GetDatabaseConnection().Model(&entities.User{}).Where("username = ? OR email = ?", username, username).First(&user).Error
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid user credentials")
	}

	return user, nil
}

func (selfC userRepository) CreateUser(user *entities.User) (*entities.User, error) {

	err := selfC.db.GetDatabaseConnection().Create(user).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}

func (selfC userRepository) GetUserByUsernameOrEmail(username, email string) (*entities.User, error) {
	var user entities.User

	err := selfC.db.GetDatabaseConnection().Where("username = ? OR email = ?", username, email).First(&user).Error

	if err != nil {
		return nil, errors.New("could not execute query")
	}

	return &user, nil
}

package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ropc-service/conf"
	"ropc-service/model/entities"
)

type ClientRepository interface {
	GetClient(clientId string) (*entities.Client, error)
}

type ClientRepositoryImpl struct {
	db *gorm.DB
}

func NewClientRepository() *ClientRepositoryImpl {
	return &ClientRepositoryImpl{db: conf.DB}
}

func (c ClientRepositoryImpl) GetClient(clientId string) (*entities.Client, error) {

	var client entities.Client

	err := c.db.Model(&entities.Client{}).Where("client_id = ?", clientId).First(&client).Error
	if err != nil {
		return nil, errors.New("invalid client")
	}

	return &client, nil
}

package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ropc-service/model"
)

type ClientRepository interface {
	GetClient(clientId string) (*model.Client, error)
}

type ClientRepositoryImpl struct {
	db *gorm.DB
}

func NewClientRepository() *ClientRepositoryImpl {
	return &ClientRepositoryImpl{db: DatabaseConfig.db}
}

func (c ClientRepositoryImpl) GetClient(clientId string) (*model.Client, error) {

	var client model.Client

	err := c.db.Debug().Model(&model.Client{}).Where("client_id = ?", clientId).First(&client).Error
	if err != nil {
		return nil, errors.New("invalid client")
	}

	return &client, nil
}

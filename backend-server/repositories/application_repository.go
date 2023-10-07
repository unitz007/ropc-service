package repositories

import (
	"errors"
	"ropc-service/conf"
	"ropc-service/model"
	"strings"

	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Get(clientId string) (*model.Application, error)
	GetAll() []model.Application
	Create(client *model.Application) error
}

type applicationRepository struct {
	db conf.Database[gorm.DB]
}

func NewClientRepository(db conf.Database[gorm.DB]) ApplicationRepository {
	return &applicationRepository{db: db}
}

func (c applicationRepository) Get(clientId string) (*model.Application, error) {

	var client model.Application

	err := c.db.GetDatabaseConnection().
		Model(&model.Application{}).
		Where("client_id = ?", clientId).
		First(&client).
		Error
	
	if err != nil {
		return nil, errors.New("application not found")
	}

	return &client, nil
}

func (c applicationRepository) GetAll() []model.Application {

	var clients []model.Application

	c.db.GetDatabaseConnection().Find(&clients)

	return clients
}

func (c applicationRepository) Create(client *model.Application) error {

	err := c.db.GetDatabaseConnection().Create(client).Error

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("application already exists")
		} else {
			return errors.New("could not create application. Contact administrator")
		}
	}

	return err
}

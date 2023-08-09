package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"ropc-service/conf"
	"ropc-service/model/entities"
	"strings"
)

type ClientRepository interface {
	GetClient(clientId string) (*entities.Client, error)
	GetClients() []entities.Client
	CreateClient(client *entities.Client) error
}

type clientRepository struct {
	db conf.Database[gorm.DB]
}

func NewClientRepository(db conf.Database[gorm.DB]) ClientRepository {
	return &clientRepository{db: db}
}

func (c clientRepository) GetClient(clientId string) (*entities.Client, error) {

	var client entities.Client

	err := c.db.GetDatabaseConnection().Model(&entities.Client{}).Where("client_id = ?", clientId).First(&client).Error
	if err != nil {
		return nil, errors.New("invalid client")
	}

	return &client, nil
}

func (c clientRepository) GetClients() []entities.Client {

	var clients []entities.Client

	c.db.GetDatabaseConnection().Find(&clients)

	return clients
}

func (c clientRepository) CreateClient(client *entities.Client) error {

	log.Println(client)

	err := c.db.GetDatabaseConnection().Create(client).Error

	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("client Id already exists")
		} else {
			return errors.New("could not create client")
		}
	}

	return err
}

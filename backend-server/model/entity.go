package model

import (
	"errors"

	"gorm.io/gorm"
)

type Application struct {
	gorm.Model   `json:"-"`
	ClientId     string `gorm:"index;unique;not-null"`
	ClientSecret string `gorm:"not-null;size:100" json:"-"`
}

func NewApplication(clientId, clientSecret string) (*Application, error) {

	if clientId == "" {
		return nil, errors.New("client id should not be empty")
	}

	app := &Application{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	return app, nil
}

func (a *Application) ToDTO() *ApplicationDto {
	return &ApplicationDto{ClientId: a.ClientId}
}

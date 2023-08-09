package entities

import (
	"errors"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model   `json:"-"`
	ClientId     string `gorm:"index;unique;not-null"`
	ClientSecret string `gorm:"not-null;size:100" json:"-"`
	GrantType    string `json:"-"`
}

func NewClient(clientId, clientSecret string) (*Client, error) {

	if clientId == "" {
		return nil, errors.New("client id should not be empty")
	}

	client := &Client{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	return client, nil
}

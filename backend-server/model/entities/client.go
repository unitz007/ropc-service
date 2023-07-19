package entities

import "gorm.io/gorm"

type Client struct {
	gorm.Model   `json:"-"`
	ClientId     string `gorm:"index;unique;not-null" json:"client_id"`
	ClientSecret string `gorm:"not-null;size:100"`
	GrantType    string `json:"-"`
}

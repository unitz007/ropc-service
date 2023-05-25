package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	ClientId     string `gorm:"index;unique;not-null"`
	ClientSecret string `gorm:"not-null"`
	GrantType    string
}

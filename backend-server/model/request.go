package model

type CreateApplication struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

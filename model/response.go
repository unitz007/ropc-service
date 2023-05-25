package model

type Response struct {
	Message string `json:"message,omitempty"`
	Payload any    `json:"payload"`
}

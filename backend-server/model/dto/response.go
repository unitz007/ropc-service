package dto

type Response[T any] struct {
	Message string `json:"message,omitempty"`
	Payload T      `json:"payload,omitempty"`
}

func NewResponse[T any](message string, payload T) *Response[T] {
	return &Response[T]{
		Message: message,
		Payload: payload,
	}
}

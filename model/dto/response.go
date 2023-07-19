package dto

type Response struct {
	Message string `json:"message,omitempty"`
	Payload any    `json:"payload"`
}

func NewResponse(message string, payload any) *Response {
	return &Response{
		Message: message,
		Payload: payload,
	}
}

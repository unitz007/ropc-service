package handlers

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"ropc-service/mocks"
	"ropc-service/model/entities"
	"strings"
	"testing"
)

func TestCreateClientHandler(t *testing.T) {

	var clientHandler ClientHandler

	tt := []struct {
		Name        string
		Body        io.Reader
		ShouldPanic bool
	}{
		{
			Name:        "nil body request",
			Body:        nil,
			ShouldPanic: true,
		},
		{
			Name:        "invalid json body",
			Body:        strings.NewReader(`"dhhdhdhr": "ljll"`),
			ShouldPanic: true,
		},
		{
			Name:        "valid json body",
			Body:        strings.NewReader(`{ "client_id": "test_client"}`),
			ShouldPanic: false,
		},
		{
			Name:        "valid json body with empty client_id",
			Body:        strings.NewReader(`{ "client_id": "" }`),
			ShouldPanic: true,
		},
	}

	for _, w := range tt {

		t.Run(w.Name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/clients", w.Body)
			response := httptest.NewRecorder()

			exec := func() {
				clientHandler = NewClientHandler(nil)
				clientHandler.CreateClient(response, request)
			}

			if w.ShouldPanic == true {
				assert.Panics(t, exec, "should panic due to invalid request body")
			} else {
				assert.NotPanics(t, exec, "should not panic")
			}
		})
	}

	t.Run("successful request should return 201 CREATED", func(t *testing.T) {
		clientService := new(mocks.ClientService)
		body := strings.NewReader(`{ "client_id": "test_client"}`)

		request := httptest.NewRequest(http.MethodPost, "/clients", body)
		response := httptest.NewRecorder()

		clientSecret := uuid.New().String()

		clientService.On("CreateClient", &entities.Client{ClientId: "test_client", ClientSecret: clientSecret})

		clientHandler := NewClientHandler(clientService)

		clientHandler.CreateClient(response, request)

		expected := 201
		got := response.Code

		if expected != got {
			t.Errorf("expected %v got %v", expected, got)
		}

	})
}

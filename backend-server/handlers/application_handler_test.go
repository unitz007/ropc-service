package handlers

import (
	"backend-server/mocks"
	"backend-server/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateClientHandler(t *testing.T) {

	var clientHandler ApplicationHandler

	tt := []struct {
		Name        string
		Body        io.Reader
		ShouldPanic bool
	}{
		{
			Name:        "nil body prepareRequest",
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
				clientHandler = NewApplicationHandler(nil)
				clientHandler.CreateApplication(response, request)
			}

			if w.ShouldPanic == true {
				assert.Panics(t, exec, "should panic due to invalid prepareRequest body")
			} else {
				assert.NotPanics(t, exec, "should not panic")
			}
		})
	}

	t.Run("successful prepareRequest should return 201 CREATED", func(t *testing.T) {
		applicationRepository := new(mocks.ApplicationRepository)
		body := strings.NewReader(`{ "client_id": "test_client", "client_secret": "test_secret"}`)

		request := httptest.NewRequest(http.MethodPost, "/apps", body)
		response := httptest.NewRecorder()

		applicationRepository.On("Create", &model.Application{ClientId: "test_client", ClientSecret: "test_secret"}).Return(nil)

		clientHandler := NewApplicationHandler(applicationRepository)

		clientHandler.CreateApplication(response, request)

		expected := 201
		got := response.Code

		if expected != got {
			t.Errorf("expected %v got %v", expected, got)
		}

	})
}

//func TestGenerateClientSecret(t *testing.T) {
//	t.Run("should contain client id as path variable", func(t *testing.T) {
//		urlWithoutClientId := "http://localhost:8080/apps"
//
//		request := httptest.NewRequest(http.MethodPost, urlWithoutClientId, w.Body)
//		response := httptest.NewRecorder()
//
//		repoMock := new(mocks.ApplicationRepository)
//
//		handler := NewApplicationHandler(a)
//
//	})
//}

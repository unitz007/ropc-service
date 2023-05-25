package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ropc-service/dto"
	"ropc-service/model"
	"ropc-service/services"
)

func Authentication(ctx *gin.Context) {

	var loginRequest *dto.LoginRequest

	err := ctx.BindJSON(&loginRequest)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, &model.Response{Message: "Invalid request body"})
		return
	}

	userAuthenticator := services.InstantiateUserAuthenticator()
	clientAuthenticator := services.InstantiateClientAuthenticator()

	user := &model.User{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}

	client := &model.Client{
		ClientId:     loginRequest.ClientId,
		ClientSecret: loginRequest.ClientSecret,
		GrantType:    loginRequest.GrantType,
	}

	token, err := services.InstantiateAuthenticator(userAuthenticator, clientAuthenticator).Authenticate(user, client)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &model.Response{Message: err.Error()})
		return
	}

	ctx.JSON(200, &model.Response{
		Message: "Authentication Successful",
		Payload: token,
	})
}

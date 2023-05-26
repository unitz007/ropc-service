package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/services"
)

func Authentication(ctx *gin.Context) {

	var loginRequest *dto.LoginRequest

	err := ctx.BindJSON(&loginRequest)

	if err != nil {
		log.Println(err.Error())
		ctx.JSON(400, &dto.Response{Message: "Invalid request body"})
		return
	}

	userAuthenticator := services.InstantiateUserAuthenticator()
	clientAuthenticator := services.InstantiateClientAuthenticator()

	user := &entities.User{
		Username: loginRequest.Username,
		Password: loginRequest.Password,
	}

	client := &entities.Client{
		ClientId:     loginRequest.ClientId,
		ClientSecret: loginRequest.ClientSecret,
		GrantType:    loginRequest.GrantType,
	}

	token, err := services.InstantiateAuthenticator(userAuthenticator, clientAuthenticator).Authenticate(user, client)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{Message: err.Error()})
		return
	}

	ctx.JSON(200, &dto.Response{
		Message: "Authentication Successful",
		Payload: token,
	})
}

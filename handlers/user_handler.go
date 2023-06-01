package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ropc-service/model/dto"
	"ropc-service/model/entities"
	"ropc-service/repositories"
	"ropc-service/services"
)

func CreateUser(ctx *gin.Context) {

	var user *entities.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request payload")
		return
	}

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)

	_, err = userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &dto.Response{
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	ctx.JSON(http.StatusCreated, &dto.Response{
		Message: "User created Successfully",
		Payload: nil,
	})

}

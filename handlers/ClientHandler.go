package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ropc-service/repositories"
)

func GetClients(ctx *gin.Context) {

	clientRepository := repositories.NewClientRepository()

	clients := clientRepository.GetClients()

	ctx.JSON(http.StatusOK, &clients)

}

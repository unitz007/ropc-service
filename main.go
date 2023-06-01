package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ropc-service/conf"
	"ropc-service/handlers"
	"ropc-service/utils"
)

func main() {
	environmentConfig := conf.LoadConfig() // load environment configuration
	conf.InitGormConfig(environmentConfig) // initialize database

	// set up request handlers
	requestHandlers := func(router *gin.Engine) {

		// authenticate
		router.POST("/login", handlers.Authentication)
		router.GET("/login", func(context *gin.Context) {
			context.HTML(http.StatusOK, "login.html", gin.H{})
		})

		// user
		router.POST("/users", handlers.CreateUser)
		router.GET("/register", func(context *gin.Context) {
			context.HTML(http.StatusOK, "register.html", gin.H{})
		})

		// secured resource
		router.GET("/", func(context *gin.Context) {

			accessToken := context.GetHeader("Authorization")
			if accessToken == "" {
				log.Println("No access token")
				context.HTML(http.StatusUnauthorized, "white_label.html", gin.H{})
				return
			}

			claims, err := utils.ValidateToken(accessToken)
			if err != nil {
				log.Println(err)
				context.HTML(http.StatusUnauthorized, "login.html", gin.H{})
				return
			}

			context.HTML(http.StatusOK, "index.html", gin.H{
				"username":  claims["username"],
				"client_id": claims["client_id"],
				"email":     claims["email"],
			})
		})

	}

	// load gin context
	conf.InitGinContext(environmentConfig, requestHandlers)
}

package services

import (
	"github.com/golang-jwt/jwt"
	"ropc-service/conf"
	"ropc-service/model"
	"time"
)

func GenerateToken(user *model.User, client *model.Client) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expiryInMinutes := conf.GlobalConfig.TokenExpiry

	claims["exp"] = time.Now().Add(time.Duration(expiryInMinutes) * time.Minute)
	claims["username"] = user.Username
	claims["client_id"] = client.ClientId
	claims["grant_type"] = client.GrantType

	return token.SignedString([]byte(conf.GlobalConfig.TokenSecret))
}

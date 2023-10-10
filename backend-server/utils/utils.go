package utils

import (
	"backend-server/model"
	"strings"

	"github.com/golang-jwt/jwt"
)

const bearerPrefix = "Bearer "

func GenerateToken(client *model.Application, tokenSecret string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	//claims["username"] = user.Username
	claims["client_id"] = client.ClientId
	//claims["email"] = user.Email

	return token.SignedString([]byte(tokenSecret))
}

func ValidateToken(token, tokenSecret string) (jwt.MapClaims, error) {
	token = strings.TrimPrefix(token, bearerPrefix)
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}

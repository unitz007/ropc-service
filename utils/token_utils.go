package utils

import (
	"github.com/golang-jwt/jwt"
	"ropc-service/conf"
	"ropc-service/model/entities"
)

func GenerateToken(user *entities.User, client *entities.Client) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = user.Username
	claims["client_id"] = client.ClientId
	claims["grant_type"] = client.GrantType
	claims["email"] = user.Email

	return token.SignedString([]byte(conf.GlobalConfig.TokenSecret))
}

func ValidateToken(token string) (jwt.MapClaims, error) {
	tokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.GlobalConfig.TokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(jwt.MapClaims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}

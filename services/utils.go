package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"ropc-service/model"
	"strconv"
	"time"
)

type TokenUtils struct {
	securityKey string
}

func InstantiateTokenUtil() *TokenUtils {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}
	return &TokenUtils{}
}

func (u *TokenUtils) GenerateToken(user *model.User, client *model.Client) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	expiryInMinutes, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRY"))
	if err != nil {
		return "", errors.New("could not generate token")
	}

	claims["exp"] = time.Now().Add(time.Duration(expiryInMinutes) * time.Minute)
	claims["username"] = user.Username
	claims["client_id"] = client.ClientId
	claims["grant_type"] = client.GrantType

	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort       string
	DatabasePassword string
	DatabaseUser     string
	DatabaseName     string
	TokenSecret      string
	DatabaseHost     string
	DatabasePort     string
	TokenExpiry      int
	GinMode          string
}

const (
	configFileName          = ".env"
	configFileErrorMessage  = "Could not find config file: " + configFileName
	dbHost                  = "DB_HOST"
	dbPort                  = "DB_PORT"
	dbUser                  = "DB_USER"
	serverPort              = "SERVER_PORT"
	ginMode                 = "GIN_MODE"
	tokenSecret             = "TOKEN_SECRET"
	dbName                  = "DB_NAME"
	dbPassword              = "DB_PASSWORD"
	tokenExpiry             = "TOKEN_EXPIRY"
	tokenExpiryErrorMessage = "invalid token expiry"
)

var GlobalConfig *Config

func LoadConfig() *Config {

	// load .env file
	err := godotenv.Load(configFileName)
	if err != nil {
		log.Fatal(configFileErrorMessage)
	}

	GlobalConfig = &Config{
		DatabaseHost:     os.Getenv(dbHost),
		DatabasePort:     os.Getenv(dbPort),
		DatabaseUser:     os.Getenv(dbUser),
		ServerPort:       os.Getenv(serverPort),
		GinMode:          os.Getenv(ginMode),
		TokenSecret:      os.Getenv(tokenSecret),
		DatabaseName:     os.Getenv(dbName),
		DatabasePassword: os.Getenv(dbPassword),
		TokenExpiry: func() int {
			val, err := strconv.Atoi(os.Getenv(tokenExpiry))
			if err != nil {
				log.Fatal(tokenExpiryErrorMessage)
			}
			return val
		}(),
	}

	return GlobalConfig
}

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

var GlobalConfig Config

func LoadConfig() Config {

	// load .env file
	err := godotenv.Load(".env2")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GlobalConfig = Config{
		DatabaseHost:     os.Getenv("DB_HOST"),
		DatabasePort:     os.Getenv("DB_PORT"),
		DatabaseUser:     os.Getenv("DB_USER"),
		ServerPort:       os.Getenv("SERVER_PORT"),
		GinMode:          os.Getenv("GIN_MODE"),
		TokenSecret:      os.Getenv("TOKEN_SECRET"),
		DatabaseName:     os.Getenv("DB_NAME"),
		DatabasePassword: os.Getenv("DB_PASSWORD"),
		TokenExpiry: func() int {
			val, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRY"))
			if err != nil {
				log.Fatal("Error: Could not load token expiry")
			}
			return val
		}(),
	}

	return GlobalConfig
}

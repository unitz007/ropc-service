package repositories

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"ropc-service/model"
)

type DBConfig struct {
	db *gorm.DB
}

var DatabaseConfig DBConfig

func InitGormConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	database, err := gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(&model.User{}, &model.Client{})
	if err != nil {
		log.Fatal("Could not migrate:", err.Error())
	}

	DatabaseConfig = DBConfig{db: database}
}

func (d DBConfig) PreLoadData() {

	log.Println("Preloading data")

	client := model.Client{
		ClientId:     "testClient",
		ClientSecret: "$2a$12$cUC8531lrEhDzRT2aOZr8eUkTRuP0mpyyRZ.nvKRBg9oL145RT8Lu", // raw -> clientSecret
	}

	d.db.Create(&client)

	user := model.User{
		Username: "testUser",
		Password: "$2a$12$ga6jVPJeORwhba8AmsoapemKDd1Z2CuFIi4bhXZapjNxnaHpHXcj6", // -> strongPassword
		Client:   client,
	}

	d.db.Create(&user)

}

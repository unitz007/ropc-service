package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ropc-service/model/entities"
)

var DB *gorm.DB

func InitGormConfig(globalConfig *Config) {
	DbHost := globalConfig.DatabaseHost
	DbUser := globalConfig.DatabaseUser
	DbPassword := globalConfig.DatabasePassword
	DbName := globalConfig.DatabaseName
	DbPort := globalConfig.DatabasePort

	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	database, err := gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = database

	err = DB.AutoMigrate(&entities.User{}, &entities.Client{})
	if err != nil {
		log.Fatal("Could not migrate:", err.Error())
	}

	PreLoadData()

}

func PreLoadData() {

	log.Println("Preloading data")

	client := entities.Client{
		ClientId:     "testClient",
		ClientSecret: "$2a$12$cUC8531lrEhDzRT2aOZr8eUkTRuP0mpyyRZ.nvKRBg9oL145RT8Lu", // raw -> clientSecret
	}

	DB.Create(&client)

	user := entities.User{
		Username: "testUser",
		Password: "$2a$12$ga6jVPJeORwhba8AmsoapemKDd1Z2CuFIi4bhXZapjNxnaHpHXcj6", // -> strongPassword
		Client:   client,
	}

	DB.Create(&user)

}

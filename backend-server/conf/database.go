package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ropc-service/model/entities"
)

type Database[T any] interface {
	GetDatabaseConnection() *T
}

type database struct {
	conn *gorm.DB
}

func NewDataBase(config Config) Database[gorm.DB] {
	DbHost := config.DatabaseHost()
	DbUser := config.DatabaseUser()
	DbPassword := config.DatabasePassword()
	DbName := config.DatabaseName()
	DbPort := config.DatabasePort()

	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := gorm.Open(mysql.Open(DbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&entities.User{}, &entities.Client{})
	if err != nil {
		log.Fatal("Could not migrate:", err.Error())
	}

	return &database{db}
}

func (db *database) GetDatabaseConnection() *gorm.DB {
	return db.conn
}

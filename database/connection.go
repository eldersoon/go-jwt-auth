package database

import (
	"github.com/go-auth/controller/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "root:@tcp(127.0.0.1:3306)/go_auth?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error to connect to the database: " + err.Error())
	}

	DB = connection

	connection.AutoMigrate(&model.User{})
}
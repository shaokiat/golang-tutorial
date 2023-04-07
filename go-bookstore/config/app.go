package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect() {
	dsn := "user:mypassword@tcp(localhost:3306)/bookstore_db?charset=utf8mb4&parseTime=True&loc=Local"


	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
			panic(err)
	}	
	db = d
}

func GetDB() *gorm.DB {
	return db
}
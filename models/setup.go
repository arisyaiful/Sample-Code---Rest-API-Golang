package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := "127.0.0.1"
	dbname := "db_golang"
	user := "postgres"
	password := "Asdqwe123"
	port := 5432

	dbURl := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(dbURl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&Book{}, &Category{})

	DB = database
}

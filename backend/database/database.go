package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "postgres://postgres:TUMHARA_PASSWORD@localhost:5432/expense_tracker"
	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("Failed to Connect to database")
	} else {
		DB = database
		fmt.Println("Successfully Connected To Database")
	}

}

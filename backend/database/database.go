// package database

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDataBase() {
// 	dsn := os.Getenv("DATABASE_URL")
// 	database, err := gorm.Open(
// 		postgres.Open(dsn),
// 		&gorm.Config{},
// 	)
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Failed to Connect to database")
// 	} else {
// 		DB = database
// 		fmt.Println("Successfully Connected To Database")
// 	}

// }
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found sorry")
	}
	dsn := os.Getenv("DATABASE_URL")
	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("Failed to Connect to database")
	}
	fmt.Println("Database Successfully Connected")
	DB = database

}

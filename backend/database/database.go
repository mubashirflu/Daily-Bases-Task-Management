// // package database

// // import (
// // 	"fmt"
// // 	"log"
// // 	"os"

// // 	"github.com/joho/godotenv"
// // 	"gorm.io/driver/postgres"
// // 	"gorm.io/gorm"
// // )

// // var DB *gorm.DB

// // func ConnectDataBase() {
// // 	dsn := os.Getenv("DATABASE_URL")
// // 	database, err := gorm.Open(
// // 		postgres.Open(dsn),
// // 		&gorm.Config{},
// // 	)
// // 	err := godotenv.Load()
// // 	if err != nil {
// // 		log.Fatal("Failed to Connect to database")
// // 	} else {
// // 		DB = database
// // 		fmt.Println("Successfully Connected To Database")
// // 	}

// // }
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
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("No .env file found sorry")
// 	}
// 	dsn := os.Getenv("DATABASE_URL")
// 	database, err := gorm.Open(
// 		postgres.Open(dsn),
// 		&gorm.Config{},
// 	)
// 	if err != nil {
// 		log.Fatal("Failed to Connect to database")
// 	}
// 	fmt.Println("Database Successfully Connected")
// 	DB = database

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

	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Connect PostgreSQL
	database, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database

	fmt.Println("Database Successfully Connected")

	// =========================
	// DEBUG: Check database
	// =========================

	var dbName string

	err = DB.Raw("SELECT current_database()").Scan(&dbName).Error
	if err != nil {
		log.Fatal("Could not get database name:", err)
	}

	fmt.Println("DATABASE:", dbName)

	var schema string

	err = DB.Raw("SELECT current_schema()").Scan(&schema).Error
	if err != nil {
		log.Fatal("Could not get schema:", err)
	}

	fmt.Println("SCHEMA:", schema)
}

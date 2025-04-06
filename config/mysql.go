package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("missing database connection string")
	}

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	db = database
	fmt.Println("Connected to database!")
}

func GetDatabase() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized, call ConnectDatabase() first")
	}
	return db
}

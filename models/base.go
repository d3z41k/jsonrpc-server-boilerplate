package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?", dbUser, dbPassword, dbHost, dbPort, dbName)
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"
	dsn += "&parseTime=true"
	dsn += "&loc=Local"

	fmt.Println(dsn)

	conn, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

// GetDB return a descriptor to the DB object
func GetDB() *gorm.DB {
	return db
}

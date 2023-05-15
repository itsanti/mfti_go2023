package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"pokemon-rest-api/listing"
)

type Config struct {
	DBHost         string
	DBUserName     string
	DBUserPassword string
	DBName         string
	DBPort         string
}

var config = Config{
	DBHost:         "127.0.0.1",
	DBUserName:     "postgres",
	DBUserPassword: "postgres",
	DBName:         "postgres",
	DBPort:         "5432",
}

var DB *gorm.DB

func connectDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUserName,
		config.DBUserPassword,
		config.DBName,
		config.DBPort,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")
}

func InitDB() {
	connectDB()
	DB.AutoMigrate(&listing.Pokemon{})
	fmt.Println("? Migration complete")
}

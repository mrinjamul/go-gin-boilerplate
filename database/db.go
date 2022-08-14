package database

import (
	"fmt"
	"go-gin-boilerplate/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// IsConnected returns the connection status
	IsConnected bool
	IsSQLite    bool
	DB          *gorm.DB
)

func GetDB() *gorm.DB {
	var (
		dbHost     string
		dbPort     string
		dbUser     string
		dbPassword string
		dbName     string
	)
	// Get ENV variables
	_, present := os.LookupEnv("POSTGRES_HOST")
	if present {
		dbHost = os.Getenv("POSTGRES_HOST")
	} else {
		dbHost = "localhost"
	}
	_, present = os.LookupEnv("POSTGRES_PORT")
	if present {
		dbPort = os.Getenv("POSTGRES_PORT")
	} else {
		dbPort = "5432"
	}
	_, present = os.LookupEnv("POSTGRES_USER")
	if present {
		dbUser = os.Getenv("POSTGRES_USER")
	} else {
		dbUser = "postgres"
	}
	_, present = os.LookupEnv("POSTGRES_PASSWORD")
	if present {
		dbPassword = os.Getenv("POSTGRES_PASSWORD")
	} else {
		dbPassword = "postgres"
	}
	_, present = os.LookupEnv("POSTGRES_DB")
	if present {
		dbName = os.Getenv("POSTGRES_DB")
	} else {
		dbName = "postgres"
	}
	// Connect to database
	if DB == nil {
		if dbHost == "" {
			fmt.Println("Environment variable DB_HOST is null.")
			return nil
		}
		if dbName == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil
		}
		if dbUser == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil
		}
		if dbPassword == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil
		}

		if dbPort == "" {
			dbPort = "5432"
		}
	}

	// Connect to DB
	dest := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err := gorm.Open(postgres.Open(dest), &gorm.Config{})

	if err == nil {
		IsConnected = true
	} else {
		log.Println("failed to connect database(pgsql)")
	}

	// if unable to connect to database, create sqlite database
	if !IsConnected {
		// Create sqlite connection
		DB, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
		if err == nil {
			IsConnected = true
			IsSQLite = true
		} else {
			log.Println("failed to connect database(sqlite)")
		}
	}

	// Migrate the schema
	DB.AutoMigrate(&models.Quote{})
	return DB
}

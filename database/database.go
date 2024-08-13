package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/aliyilmazdev/personal-budget-tracker/config"
	"github.com/aliyilmazdev/personal-budget-tracker/internal/transaction"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Error parsing port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connected!")

	DB.AutoMigrate(&transaction.Transaction{})
    fmt.Println("Database Migrated")
}
package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config stores PostgreSQL connection details
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection creates a new PostgreSQL connection using GORM
func NewConnetion(cfg *Config) (*gorm.DB, error) {
	// Correct: define dsn as a string variable
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("error connecting to database: %v", err)
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}

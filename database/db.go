package database

import (
	"fmt"
	"go_assignment2/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "go_assignment2"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Default().Println("Connection to db success")

	err = migration(db)
	if err != nil {
		panic(err)
	}
	return db
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.Order{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.Item{}); err != nil {
		return err
	}
	return nil
}

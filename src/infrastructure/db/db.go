package db

import (
	"log"
	"os"

	"github.com/erickrodrigs/go-webapi-template/src/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB ...
func ConnectDB() *gorm.DB {
	dsn := os.Getenv("dsn")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.AutoMigrate(&model.Book{})

	return db
}

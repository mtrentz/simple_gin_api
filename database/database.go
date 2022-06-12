package database

import (
	"log"

	"github.com/mtrentz/simple_gin_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Error connecting to the database")
	}
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Author{})
	DB.AutoMigrate(&models.Publisher{})
}
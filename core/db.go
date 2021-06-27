package core

import (
	"os"
	"fmt"
	"log"
	"time"
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
)

type Company struct {
	ID        uint           `gorm:"primaryKey"`
	Name string
	Zipcode string
	Website string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func GetDb() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
						os.Getenv("POSTGRES_HOST"),
						os.Getenv("POSTGRES_USER"),
						os.Getenv("POSTGRES_PASSWORD"),
						os.Getenv("POSTGRES_DB_NAME"),
						os.Getenv("POSTGRES_PORT"))

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func InitializeDb() {
	db, err := GetDb()

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Company{})
}
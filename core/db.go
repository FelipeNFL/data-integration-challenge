package core

import (
	"os"
	"fmt"
	"log"
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"
)

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

func CleanDb() {
	db, err := GetDb()

	if err != nil {
		log.Fatal(err)
	}

	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Company{})
}
package core

import (
	"time"
	"gorm.io/gorm"
)

type Company struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Zipcode string `json:"zipcode"`
	Website string `json:"website"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
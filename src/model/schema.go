package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:""`
}

type Books struct {
	gorm.Model
	id      uint   `gorm:"primaKey"`
	name    string `gorm:"not null"`
	edition string `gorm:""`
}

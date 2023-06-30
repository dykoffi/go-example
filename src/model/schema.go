package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Name string `gorm:""`
	Mail string `gorm:"<-:false"`
	Name string `gorm:""`
	Age  uint64 `gorm:"check:age > 15"`
}

type Books struct {
	gorm.Model
	Id      uint   `gorm:"primaKey"`
	Name    string `gorm:"not null"`
	Edition string `gorm:""`
}

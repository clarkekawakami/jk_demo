package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int
	Name         string
	Email        string
	Age          int
	Appointments []Appointment `gorm:"foreignKey:UserID"`
}

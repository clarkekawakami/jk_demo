package models

import (
	"gorm.io/gorm"
)

type Facility struct {
	gorm.Model
	ID           int
	Name         string
	Description  string
	Address      string
	City         string
	State        string
	Zip          string
	Telephone    string
	Email        string
	Resources    []Resource    `gorm:"foreignKey:FacilityID"`
	Appointments []Appointment `gorm:"foreignKey:FacilityID"`
}

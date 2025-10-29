package models

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	ID           int
	Name         string
	Description  string
	FacilityID   uint
	Appointments []Appointment `gorm:"foreignKey:ResourceID"`
}

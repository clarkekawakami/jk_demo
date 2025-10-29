package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	ID         int
	Name       string
	FacilityID uint
	ResourceID uint
	UserID     uint
	Status     string
	Appt_Date  string
	Appt_Time  string
}

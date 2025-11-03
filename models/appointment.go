package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	ID         int
	Name       string
	FacilityID int
	ResourceID int
	UserID     int
	Status     string
	Appt_Date  string
	Appt_Time  string
}

type PostSearchRequest struct {
	FacilityID   int
	ReqDate      string
	ReqTimeRange string
}

type PostSearchResponse struct {
	Status string `json:"stat"`
	Result bool   `json:"result"`
}

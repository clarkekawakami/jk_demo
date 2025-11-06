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
	FacilityID int
	Appt_Date  string
	Appt_Time  string
}

type NewAppointmentView struct {
	FacilityDesc    string
	FacilityAddress string
	FacilityCity    string
	FacilityState   string
	FacilityZip     string
	FacilityPhone   string
	ResourceDesc    string
	Appt_Date       string
	Appt_Time       string
}

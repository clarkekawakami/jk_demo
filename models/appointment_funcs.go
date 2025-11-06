package models

import (
	"gorm.io/gorm"
)

// Appointment entity functions

// create a Appointment
func CreateAppointment(db *gorm.DB, Appointment *Appointment) (err error) {
	// fmt.Println(Appointment)
	err = db.Create(Appointment).Error
	if err != nil {
		return err
	}
	return nil
}

// get Appointment
func GetAppointments(db *gorm.DB, Appointment *[]Appointment) (err error) {
	err = db.Find(Appointment).Error
	if err != nil {
		return err
	}
	return nil
}

// get Appointment by id
func GetAppointment(db *gorm.DB, Appointment *Appointment, id int) (err error) {
	err = db.Where("id = ?", id).First(Appointment).Error
	if err != nil {
		return err
	}
	return nil
}

// update Appointment
func UpdateAppointment(db *gorm.DB, Appointment *Appointment) (err error) {
	db.Save(Appointment)
	return nil
}

// delete Appointment
func DeleteAppointment(db *gorm.DB, Appointment *Appointment, id int) (err error) {
	db.Where("id = ?", id).Delete(Appointment)
	return nil
}

// truncate Appointments table
func TruncateAppointments(db *gorm.DB) (err error) {
	err = db.Where("id > 0").Delete(&Appointment{}).Error
	if err != nil {
		return err
	}
	return nil
}

// get Appointment
// db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&usersInDateRange)
func GetDateFacilityAppointments(db *gorm.DB, Appointment *[]Appointment, reqDate string, facility_id int, start_time string, end_time string) (err error) {
	err = db.Where("appt_date = ? and facility_id = ? and appt_time between ? and ?", reqDate, facility_id, start_time, end_time).Order("appt_time ASC").Find(Appointment).Error
	if err != nil {
		return err
	}
	return nil
}

// get new appointment view
func GetNewAppointmentView(db *gorm.DB, newApptView *NewAppointmentView, appointment_id int) (err error) {
	var results []NewAppointmentView
	err = db.Model(&Appointment{}).Select("facilities.description as facility_desc, facilities.address as facility_address, facilities.city as facility_city, facilities.state as facility_state, facilities.zip as facility_zip, facilities.telephone as facility_phone, resources.description as resource_desc, appointments.appt_date, appointments.appt_time").
		Joins("join facilities on facilities.id = appointments.facility_id").
		Joins("join resources on resources.id = appointments.resource_id").
		Where("appointments.id = ?", appointment_id).
		Find(&results).Error
	if err != nil {
		return err
	}
	*newApptView = results[0]
	return nil

}

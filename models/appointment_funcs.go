package models

import "gorm.io/gorm"

// Appointment entity functions

// create a Appointment
func CreateAppointment(db *gorm.DB, Appointment *Appointment) (err error) {
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

package models

import (
	"gorm.io/gorm"
)

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

// // reseed Appointments table
// func ReseedAppointments(db *gorm.DB) (err error) {
// 	var newStatus string
// 	// get today's date
// 	now := time.Now()
// 	dateFormat := "2006-01-02"
// 	apptCounter := 1

// 	apptTimes := []string{"07:00:00",
// 		"07:30:00",
// 		"08:00:00",
// 		"08:30:00",
// 		"09:00:00",
// 		"09:30:00",
// 		"10:00:00",
// 		"10:30:00",
// 		"11:00:00",
// 		"11:30:00",
// 		"12:00:00",
// 		"12:30:00",
// 		"13:00:00",
// 		"13:30:00",
// 		"14:00:00",
// 		"14:30:00",
// 		"15:00:00",
// 		"15:30:00",
// 		"16:00:00",
// 		"16:30:00",
// 		"17:00:00",
// 		"17:30:00",
// 		"18:00:00",
// 		"18:30:00",
// 	}

// 	//first delete all existing records
// 	err = db.Exec("DELETE FROM appointments").Error
// 	if err != nil {
// 		return err
// 	}
// 	// get all facilities
// 	var facilities []Facility
// 	err = db.Find(&facilities).Error
// 	if err != nil {
// 		return err
// 	}
// 	// for each facility
// 	for _, facility := range facilities {
// 		// get all resources for the facility
// 		var resources []Resource
// 		err = db.Where("facility_id = ?", facility.ID).Find(&resources).Error
// 		if err != nil {
// 			return err
// 		}
// 		// for each resource
// 		for _, resource := range resources {
// 			// for the next 30 days
// 			for i := 0; i < 30; i++ {
// 				apptDate := now.AddDate(0, 0, i).Format(dateFormat)
// 				// only create appointments on weekdays
// 				weekday := now.AddDate(0, 0, i).Weekday()
// 				if weekday == time.Saturday || weekday == time.Sunday {
// 					continue
// 				} else {
// 					// create 3 appointments per day at 10:00, 13:00, 15:00
// 					for _, apptTime := range apptTimes {
// 						// rand.Seed(time.Now().UnixNano())
// 						if rand.Float32() < 0.3 {
// 							newStatus = "Scheduled"
// 						} else {
// 							newStatus = "Open"
// 						}

// 						appointment := Appointment{
// 							FacilityID: facility.ID,
// 							ResourceID: resource.ID,
// 							Appt_Date:   apptDate,
// 							Appt_Time:   apptTime,
// 							UserID:     2,
// 							Status:     "Booked",
// 							Name:       fmt.Sprintf("APT-%d", apptCounter),
// 						}
// 						err1 := models.CreateAppointment(repository.Db, &appointment)
// 						if err1 != nil {
// 							c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
// 							return
// 						}

// 						apptCounter++
// 					}

// 				}
// 			}
// 		}
// 	}

// 	return nil
// }

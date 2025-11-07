package controllers

import (
	"jk_demo/models"
	"math/rand/v2"
	"time"

	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// get html appointments page
func (repository *UserRepo) GetAppointmentsPage(c *gin.Context) {
	var appointment []models.Appointment

	err := models.GetAppointments(repository.Db, &appointment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Printf("%+v\n", appointment)
	// fmt.Println("appointment::::::", appointment)

	c.HTML(http.StatusOK,
		"appointments_page.html",
		gin.H{
			"appointments": appointment,
			"title":        "Appointments Table",
			"subtitle":     "List",
			"active":       "appointment",
		},
	)

}

// get html appointments list
func (repository *UserRepo) GetAppointmentsList(c *gin.Context) {
	var appointment []models.Appointment

	err := models.GetAppointments(repository.Db, &appointment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("%+v\n", appointment)
	// c.JSON(http.StatusOK, appointment)
	c.HTML(http.StatusOK,
		"appointment_list.html",
		gin.H{
			"appointments": appointment,
			"subtitle":     "List",
		},
	)
}

// create Appointment
func (repository *UserRepo) CreateAppointment(c *gin.Context) {
	var appointment models.Appointment

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&appointment); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("appointment after bind::::::", appointment)
	appointment.Status = "Scheduled"
	fmt.Println("appointment after status::::::", appointment)

	err := models.CreateAppointment(repository.Db, &appointment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		// return
	} else {
		fmt.Println("appointment created with ID::::::", appointment.ID)
	}
	newApptId := appointment.ID
	fmt.Println("newApptId::::::", newApptId)
	fmt.Println("output format:::::: HTML")

	// now get the appointment list
	var appointments []models.Appointment

	err1 := models.GetAppointments(repository.Db, &appointments)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"appointment_list.html",
		gin.H{
			"appointments": appointments,
			"subtitle":     "List",
		},
	)
}

// create Appointment
func (repository *UserRepo) CreateSelectedAppointment(c *gin.Context) {
	var appointment models.Appointment

	if c.Param("output") == "html" {
		fmt.Println("output format:::::: HTML")
	} else {
		fmt.Println("output format:::::: JSON")
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&appointment); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("appointment after bind::::::", appointment)
	appointment.Status = "Scheduled"
	appointment.UserID = 2
	appointment.FacilityID, _ = strconv.Atoi(c.Param("fac_id"))
	appointment.Appt_Date = c.Param("appt_date")
	appointment.Appt_Time = c.Param("appt_time")
	fmt.Println("appointment after assignments b4 resources::::::", appointment)

	// need to assign a resource that is not already booked for that date/time
	var resources []models.Resource
	err := models.GetFacilityResources(repository.Db, &resources, appointment.FacilityID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println("resources for facility::::::", resources)

	// get existing appointments for that date, facility and time
	var existingAppointments []models.Appointment
	err = models.GetDateFacilityAppointments(repository.Db, &existingAppointments, appointment.Appt_Date, appointment.FacilityID, appointment.Appt_Time, appointment.Appt_Time)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println("existingAppointments for date/time::::::", existingAppointments)
	fmt.Println("first resource", resources[0].ID)

	// find first resource that is not already booked
	appointment.ResourceID = resources[0].ID

OuterLoop:
	for _, resource := range resources {
		// isBooked := false
		for _, existingAppt := range existingAppointments {
			if existingAppt.ResourceID != resource.ID {
				appointment.ResourceID = resource.ID
				break OuterLoop
			}
		}
	}

	fmt.Println("appointment after assignments w/ resources::::::", appointment)

	err = models.CreateAppointment(repository.Db, &appointment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		// return
	} else {
		fmt.Println("appointment created with ID::::::", appointment.ID)
	}
	newApptId := appointment.ID
	fmt.Println("newApptId::::::", newApptId)

	// now if sending json to the api caller, return the new appointment record as json
	if c.Param("output") == "json" {
		var newAppointment models.Appointment
		err := models.GetAppointment(repository.Db, &newAppointment, newApptId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			// return
		}
		// now get new appointment details for view
		var newAppointmentView models.NewAppointmentView
		err = models.GetNewAppointmentView(repository.Db, &newAppointmentView, newApptId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		fmt.Printf("newAppointmentView:::::: %+v\n", newAppointmentView)

		c.JSON(http.StatusOK, newAppointmentView)
		return
	} else {
		fmt.Println("output format:::::: HTML")

		// now get the appointment list
		var appointments []models.Appointment

		err1 := models.GetAppointments(repository.Db, &appointments)
		if err1 != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
			return
		}
		c.HTML(http.StatusOK,
			"appointment_list.html",
			gin.H{
				"appointments": appointments,
				"subtitle":     "List",
			},
		)
	}
}

// get html appointments form
func (repository *UserRepo) GetAppointmentForm(c *gin.Context) {
	var appointment models.Appointment
	subtitle := "Edit Record"
	action := "put"

	if c.Param("id") != "new" {

		id, _ := strconv.Atoi(c.Param("id"))
		err := models.GetAppointment(repository.Db, &appointment, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	} else {
		subtitle = "New Record"
		action = "post"
	}
	c.HTML(http.StatusOK,
		"appointment_form.html",
		gin.H{
			"appointment": appointment,
			"subtitle":    subtitle,
			"action":      action,
		},
	)

}

// update Appointment
func (repository *UserRepo) UpdateAppointment(c *gin.Context) {
	var appointment models.Appointment

	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetAppointment(repository.Db, &appointment, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&appointment); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateAppointment(repository.Db, &appointment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Appointment list
	var appointments []models.Appointment

	err1 := models.GetAppointments(repository.Db, &appointments)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"appointment_list.html",
		gin.H{
			"appointments": appointments,
			"subtitle":     "List",
		},
	)
}

// delete Appointment
func (repository *UserRepo) DeleteAppointment(c *gin.Context) {
	var appointment models.Appointment
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteAppointment(repository.Db, &appointment, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Appointment list
	var appointments []models.Appointment

	err1 := models.GetAppointments(repository.Db, &appointments)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"appointment_list.html",
		gin.H{
			"appointments": appointments,
			"subtitle":     "List",
		},
	)
}

// get open timeslots for appointment
func (repository *UserRepo) GetAvailablePage(c *gin.Context) {
	// inputs are appt_date and facility_id and time_ranges (early morning, late morning, early afternoon, late afternoon & evening)
	// 1. get time slots for the time range
	// 2. get existing appointments for that date and facility in time range
	// 3. compare and return available time slots
	// 4. for each available time slot, get resource is not already booked for that time slot
	var appointment []models.Appointment

	// err := models.GetDateFacilityAppointments(repository.Db, &appointment, "2025-11-03", 1, "07:00:00", "12:00:00")
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	// fmt.Printf("%+v\n", appointment)
	// c.JSON(http.StatusOK, appointment)
	c.HTML(http.StatusOK,
		"available_appointments_page.html",
		gin.H{
			"appointments": appointment,
			"title":        "Appointments Table",
			"subtitle":     "Available Slots Search",
			"active":       "appointment",
		},
	)

}

// post search for open appointment slots
func (repository *UserRepo) SearchForOpen(c *gin.Context) {
	var searchFor models.PostSearchRequest

	if c.Param("output") == "html" {
		fmt.Println("output format:::::: HTML")
	} else {
		fmt.Println("output format:::::: JSON")
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&searchFor); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("searchFor facilityId::::::", searchFor.FacilityID)
	fmt.Println("searchFor request date::::::", searchFor.ReqDate)
	fmt.Println("searchFor time range::::::", searchFor.ReqTimeRange)

	// determine time range
	var start_time string
	var end_time string
	var slotTimes []string

	switch searchFor.ReqTimeRange {
	case "early_morning":
		// 07:00 to 09:00
		start_time = "07:00:00"
		end_time = "09:00:00"
		slotTimes = []string{"07:00:00",
			"07:30:00",
			"08:00:00",
			"08:30:00",
			"09:00:00",
		}
	case "late_morning":
		// 09:00 to 12:00
		start_time = "09:00:00"
		end_time = "12:00:00"
		slotTimes = []string{"09:00:00",
			"09:30:00",
			"10:00:00",
			"10:30:00",
			"11:00:00",
			"11:30:00",
			"12:00:00",
		}
	case "early_afternoon":
		// 12:00 to 15:00
		start_time = "12:00:00"
		end_time = "15:00:00"
		slotTimes = []string{"12:00:00",
			"12:30:00",
			"13:00:00",
			"13:30:00",
			"14:00:00",
			"14:30:00",
			"15:00:00",
		}
	case "late_afternoon":
		// 15:00 to 17:00
		start_time = "15:00:00"
		end_time = "17:00:00"
		slotTimes = []string{"15:00:00",
			"15:30:00",
			"16:00:00",
			"16:30:00",
			"17:00:00",
		}
	case "evening":
		// default to evening
		start_time = "17:00:00"
		end_time = "18:30:00"
		slotTimes = []string{"17:00:00",
			"17:30:00",
			"18:00:00",
			"18:30:00",
		}
	default:
		// default to all day 07:00 to 18:30
		start_time = "07:00:00"
		end_time = "19:00:00"
		slotTimes = []string{"07:00:00",
			"07:30:00",
			"08:00:00",
			"08:30:00",
			"09:00:00",
			"09:30:00",
			"10:00:00",
			"10:30:00",
			"11:00:00",
			"11:30:00",
			"12:00:00",
			"12:30:00",
			"13:00:00",
			"13:30:00",
			"14:00:00",
			"14:30:00",
			"15:00:00",
			"15:30:00",
			"16:00:00",
			"16:30:00",
			"17:00:00",
			"17:30:00",
			"18:00:00",
			"18:30:00",
		}
	}

	fmt.Println("determined time range::::::", start_time, " to ", end_time)
	fmt.Println("determined slot times::::::", slotTimes)

	// get existing appointments for that date and facility in time range
	var appointments []models.Appointment
	err := models.GetDateFacilityAppointments(repository.Db, &appointments, searchFor.ReqDate, searchFor.FacilityID, start_time, end_time)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// foreach slot time
	// query to get # of appointments for that date, facility and time
	// query to get # of resources for that facility
	// if appts < resources then add to available list

	var availableAppointments []string

	for _, slotTime := range slotTimes {
		var apptCount int64
		err := repository.Db.Model(&models.Appointment{}).Where("appt_date = ? and facility_id = ? and appt_time = ?", searchFor.ReqDate, searchFor.FacilityID, slotTime).Count(&apptCount).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var resourceCount int64
		err = repository.Db.Model(&models.Resource{}).Where("facility_id = ?", searchFor.FacilityID).Count(&resourceCount).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		fmt.Println("slotTime::::::", slotTime, " apptCount::::::", apptCount, " resourceCount::::::", resourceCount)
		if apptCount < resourceCount {
			availableAppointments = append(availableAppointments, slotTime)
		}
	}

	fmt.Println("availableAppointments::::::", availableAppointments)

	// build appointment structs for available times
	var appointmentsResult []models.PostSearchResponse
	for _, availTime := range availableAppointments {
		appt := models.PostSearchResponse{
			FacilityID: searchFor.FacilityID,
			Appt_Date:  searchFor.ReqDate,
			Appt_Time:  availTime,
		}
		appointmentsResult = append(appointmentsResult, appt)
	}

	fmt.Println("appointmentsResult::::::", appointmentsResult)

	// now get the appointment list
	// var appointments []models.Appointment

	// err1 := models.GetAppointments(repository.Db, &appointments)
	// if err1 != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
	// 	return
	// }

	if c.Param("output") == "html" {
		c.HTML(http.StatusOK,
			"available_list_body1.html",
			gin.H{
				"appointments": appointmentsResult,
				"subtitle":     "Available List",
			},
		)
	} else {
		c.JSON(http.StatusOK, appointmentsResult)
	}
}

// reseed the appointments table
func (repository *UserRepo) ReseedAppointments(c *gin.Context) {
	fmt.Println("Reseeding appointments...")
	// reseed Appointments table
	// func ReseedAppointments(db *gorm.DB) (err error) {

	// get today's date
	now := time.Now()
	dateFormat := "2006-01-02"

	apptCounter := 1
	apptProb := 0.3

	apptTimes := []string{"07:00:00",
		"07:30:00",
		"08:00:00",
		"08:30:00",
		"09:00:00",
		"09:30:00",
		"10:00:00",
		"10:30:00",
		"11:00:00",
		"11:30:00",
		"12:00:00",
		"12:30:00",
		"13:00:00",
		"13:30:00",
		"14:00:00",
		"14:30:00",
		"15:00:00",
		"15:30:00",
		"16:00:00",
		"16:30:00",
		"17:00:00",
		"17:30:00",
		"18:00:00",
		"18:30:00",
	}
	// fmt.Println("truncating...")
	//first delete all existing records
	err := models.TruncateAppointments(repository.Db)
	// err := db.Exec("DELETE FROM appointments").Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Println("getting facilities...")
	// get all facilities
	var facilities []models.Facility
	err = models.GetFacilitys(repository.Db, &facilities)
	// err = db.Find(&facilities).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// for each facility
	// fmt.Println("foreach facility..")
	for _, facility := range facilities {
		// get all resources for the facility
		fmt.Println("getting resources...")
		var resources []models.Resource
		err = models.GetFacilityResources(repository.Db, &resources, facility.ID)
		// err = db.Where("facility_id = ?", facility.ID).Find(&resources).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		// for each resource
		// fmt.Println("foreach resource..")
		for _, resource := range resources {
			// for the next 30 days
			for i := 0; i < 30; i++ {
				apptDate := now.AddDate(0, 0, i).Format(dateFormat)
				// only create appointments on weekdays
				weekday := now.AddDate(0, 0, i).Weekday()
				// set  probability of creating an appointment
				if i > 20 {
					apptProb = 0.1
				} else if i > 10 {
					apptProb = 0.2
				} else if i > 5 {
					apptProb = 0.3
				} else {
					apptProb = 0.4
				}
				if weekday == time.Saturday || weekday == time.Sunday {
					fmt.Println("skipping weekend...", apptDate)
					continue
				} else {
					fmt.Println("creating appts for date...", apptDate)
					for _, apptTime := range apptTimes {
						// rand.Seed(time.Now().UnixNano())
						if rand.Float64() < apptProb {
							appointment := models.Appointment{
								FacilityID: facility.ID,
								ResourceID: resource.ID,
								Appt_Date:  apptDate,
								Appt_Time:  apptTime,
								UserID:     2,
								Status:     "Scheduled",
								Name:       fmt.Sprintf("APT-%d", apptCounter),
							}
							err1 := models.CreateAppointment(repository.Db, &appointment)
							if err1 != nil {
								c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
								return
							}
							apptCounter++
						}
					}

				}
			}
		}
		fmt.Println("done - recs created:", apptCounter)
	}

	// now get the Appointment list
	var appointments []models.Appointment

	err1 := models.GetAppointments(repository.Db, &appointments)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"appointment_list.html",
		gin.H{
			"appointments": appointments,
			"subtitle":     "List",
		},
	)
}

package controllers

import (
	
	"jk_demo/models"

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
			"appointments":     appointment,
			"title":    "Appointments Table",
			"subtitle": "List",
			"active":   "appointment",
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
			"appointments":     appointment,
			"subtitle": "List",
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
	fmt.Println("appointment after bind::::::", *&appointment)

	err := models.CreateAppointment(repository.Db, &appointment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

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
			"appointments":     appointments,
			"subtitle": "List",
		},
	)
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
			"appointment":      appointment,
			"subtitle": subtitle,
			"action":   action,
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
			"appointments":     appointments,
			"subtitle": "List",
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
			"appointments":     appointments,
			"subtitle": "List",
		},
	)
}

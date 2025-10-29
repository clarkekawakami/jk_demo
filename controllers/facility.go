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

// get html facilities page
func (repository *UserRepo) GetFacilitysPage(c *gin.Context) {
	var facility []models.Facility

	err := models.GetFacilitys(repository.Db, &facility)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Printf("%+v\n", facility)
	// fmt.Println("facility::::::", facility)

	c.HTML(http.StatusOK,
		"facilities_page.html",
		gin.H{
			"facilities": facility,
			"title":      "Facilities Table",
			"subtitle":   "List",
			"active":     "facility",
		},
	)

}

// get html facilities list
func (repository *UserRepo) GetFacilitysList(c *gin.Context) {
	var facility []models.Facility

	err := models.GetFacilitys(repository.Db, &facility)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("%+v\n", facility)
	// c.JSON(http.StatusOK, facility)
	c.HTML(http.StatusOK,
		"facility_list.html",
		gin.H{
			"facilities": facility,
			"subtitle":   "List",
		},
	)
}

// create Facility
func (repository *UserRepo) CreateFacility(c *gin.Context) {
	var facility models.Facility

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&facility); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("facility after bind::::::", *&facility)

	err := models.CreateFacility(repository.Db, &facility)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// now get the facility list
	var facilities []models.Facility

	err1 := models.GetFacilitys(repository.Db, &facilities)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"facility_list.html",
		gin.H{
			"facilities": facilities,
			"subtitle":   "List",
		},
	)
}

// get html facilities form
func (repository *UserRepo) GetFacilityForm(c *gin.Context) {
	var facility models.Facility
	subtitle := "Edit Record"
	action := "put"

	if c.Param("id") != "new" {

		id, _ := strconv.Atoi(c.Param("id"))
		err := models.GetFacility(repository.Db, &facility, id)
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
		"facility_form.html",
		gin.H{
			"facility": facility,
			"subtitle": subtitle,
			"action":   action,
		},
	)

}

// update Facility
func (repository *UserRepo) UpdateFacility(c *gin.Context) {
	var facility models.Facility

	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetFacility(repository.Db, &facility, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&facility); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateFacility(repository.Db, &facility)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Facility list
	var facilities []models.Facility

	err1 := models.GetFacilitys(repository.Db, &facilities)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"facility_list.html",
		gin.H{
			"facilities": facilities,
			"subtitle":   "List",
		},
	)
}

// delete Facility
func (repository *UserRepo) DeleteFacility(c *gin.Context) {
	var facility models.Facility
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteFacility(repository.Db, &facility, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Facility list
	var facilities []models.Facility

	err1 := models.GetFacilitys(repository.Db, &facilities)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"facility_list.html",
		gin.H{
			"facilities": facilities,
			"subtitle":   "List",
		},
	)
}

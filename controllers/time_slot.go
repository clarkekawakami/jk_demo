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

// get html time_slots page
func (repository *UserRepo) GetTime_slotsPage(c *gin.Context) {
	var time_slot []models.Time_slot

	err := models.GetTime_slots(repository.Db, &time_slot)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Printf("%+v\n", time_slot)
	// fmt.Println("time_slot::::::", time_slot)

	c.HTML(http.StatusOK,
		"time_slots_page.html",
		gin.H{
			"time_slots": time_slot,
			"title":      "Time_slots Table",
			"subtitle":   "List",
			"active":     "time_slot",
		},
	)

}

// get html time_slots list
func (repository *UserRepo) GetTime_slotsList(c *gin.Context) {
	var time_slot []models.Time_slot

	err := models.GetTime_slots(repository.Db, &time_slot)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("%+v\n", time_slot)
	// c.JSON(http.StatusOK, time_slot)
	c.HTML(http.StatusOK,
		"time_slot_list.html",
		gin.H{
			"time_slots": time_slot,
			"subtitle":   "List",
		},
	)
}

// create Time_slot
func (repository *UserRepo) CreateTime_slot(c *gin.Context) {
	var time_slot models.Time_slot

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&time_slot); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("time_slot after bind::::::", time_slot)

	err := models.CreateTime_slot(repository.Db, &time_slot)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// now get the time_slot list
	var time_slots []models.Time_slot

	err1 := models.GetTime_slots(repository.Db, &time_slots)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"time_slot_list.html",
		gin.H{
			"time_slots": time_slots,
			"subtitle":   "List",
		},
	)
}

// get html time_slots form
func (repository *UserRepo) GetTime_slotForm(c *gin.Context) {
	var time_slot models.Time_slot
	subtitle := "Edit Record"
	action := "put"

	if c.Param("id") != "new" {

		id, _ := strconv.Atoi(c.Param("id"))
		err := models.GetTime_slot(repository.Db, &time_slot, id)
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
		"time_slot_form.html",
		gin.H{
			"time_slot": time_slot,
			"subtitle":  subtitle,
			"action":    action,
		},
	)

}

// update Time_slot
func (repository *UserRepo) UpdateTime_slot(c *gin.Context) {
	var time_slot models.Time_slot

	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetTime_slot(repository.Db, &time_slot, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&time_slot); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateTime_slot(repository.Db, &time_slot)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Time_slot list
	var time_slots []models.Time_slot

	err1 := models.GetTime_slots(repository.Db, &time_slots)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"time_slot_list.html",
		gin.H{
			"time_slots": time_slots,
			"subtitle":   "List",
		},
	)
}

// delete Time_slot
func (repository *UserRepo) DeleteTime_slot(c *gin.Context) {
	var time_slot models.Time_slot
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteTime_slot(repository.Db, &time_slot, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Time_slot list
	var time_slots []models.Time_slot

	err1 := models.GetTime_slots(repository.Db, &time_slots)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"time_slot_list.html",
		gin.H{
			"time_slots": time_slots,
			"subtitle":   "List",
		},
	)
}

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


// get html resources page
func (repository *UserRepo) GetResourcesPage(c *gin.Context) {
	var resource []models.Resource

	err := models.GetResources(repository.Db, &resource)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Printf("%+v\n", resource)
	// fmt.Println("resource::::::", resource)

	c.HTML(http.StatusOK,
		"resources_page.html",
		gin.H{
			"resources":     resource,
			"title":    "Resources Table",
			"subtitle": "List",
			"active":   "resource",
		},
	)

}

// get html resources list
func (repository *UserRepo) GetResourcesList(c *gin.Context) {
	var resource []models.Resource

	err := models.GetResources(repository.Db, &resource)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("%+v\n", resource)
	// c.JSON(http.StatusOK, resource)
	c.HTML(http.StatusOK,
		"resource_list.html",
		gin.H{
			"resources":     resource,
			"subtitle": "List",
		},
	)
}

// create Resource
func (repository *UserRepo) CreateResource(c *gin.Context) {
	var resource models.Resource

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&resource); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("resource after bind::::::", *&resource)

	err := models.CreateResource(repository.Db, &resource)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// now get the resource list
	var resources []models.Resource

	err1 := models.GetResources(repository.Db, &resources)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"resource_list.html",
		gin.H{
			"resources":     resources,
			"subtitle": "List",
		},
	)
}

// get html resources form
func (repository *UserRepo) GetResourceForm(c *gin.Context) {
	var resource models.Resource
	subtitle := "Edit Record"
	action := "put"

	if c.Param("id") != "new" {

		id, _ := strconv.Atoi(c.Param("id"))
		err := models.GetResource(repository.Db, &resource, id)
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
		"resource_form.html",
		gin.H{
			"resource":      resource,
			"subtitle": subtitle,
			"action":   action,
		},
	)

}

// update Resource
func (repository *UserRepo) UpdateResource(c *gin.Context) {
	var resource models.Resource

	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetResource(repository.Db, &resource, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&resource); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateResource(repository.Db, &resource)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Resource list
	var resources []models.Resource

	err1 := models.GetResources(repository.Db, &resources)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"resource_list.html",
		gin.H{
			"resources":     resources,
			"subtitle": "List",
		},
	)
}

// delete Resource
func (repository *UserRepo) DeleteResource(c *gin.Context) {
	var resource models.Resource
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteResource(repository.Db, &resource, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the Resource list
	var resources []models.Resource

	err1 := models.GetResources(repository.Db, &resources)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"resource_list.html",
		gin.H{
			"resources":     resources,
			"subtitle": "List",
		},
	)
}

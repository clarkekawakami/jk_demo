package controllers

import (
	"jk_demo/database"

	"jk_demo/models"

	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	// db.AutoMigrate(&models.User{}, &models.Pet{}, &models.Todo{})
	db.AutoMigrate(&models.User{}, &models.Facility{}, &models.Resource{}, &models.Appointment{})
	return &UserRepo{Db: db}
}

// get html users page
func (repository *UserRepo) GetUsersPage(c *gin.Context) {
	var user []models.User

	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Printf("%+v\n", user)
	// fmt.Println("user::::::", user)

	c.HTML(http.StatusOK,
		"users_page.html",
		gin.H{
			"users":    user,
			"title":    "Users Table",
			"subtitle": "List",
			"active":   "user",
		},
	)

}

// get html users list
func (repository *UserRepo) GetUsersList(c *gin.Context) {
	var user []models.User

	err := models.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Printf("%+v\n", user)
	// c.JSON(http.StatusOK, user)
	c.HTML(http.StatusOK,
		"user_list.html",
		gin.H{
			"users":    user,
			"subtitle": "List",
		},
	)
}

// create User
func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user models.User

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("user after bind::::::", *&user)

	err := models.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// now get the user list
	var users []models.User

	err1 := models.GetUsers(repository.Db, &users)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"user_list.html",
		gin.H{
			"users":    users,
			"subtitle": "List",
		},
	)
}

// get html users form
func (repository *UserRepo) GetUserForm(c *gin.Context) {
	var user models.User
	subtitle := "Edit Record"
	action := "put"

	if c.Param("id") != "new" {

		id, _ := strconv.Atoi(c.Param("id"))
		err := models.GetUser(repository.Db, &user, id)
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
		"user_form.html",
		gin.H{
			"user":     user,
			"subtitle": subtitle,
			"action":   action,
		},
	)

}

// update User
func (repository *UserRepo) UpdateUser(c *gin.Context) {
	var user models.User

	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetUser(repository.Db, &user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("bind error:::::::::", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the User list
	var users []models.User

	err1 := models.GetUsers(repository.Db, &users)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"user_list.html",
		gin.H{
			"users":    users,
			"subtitle": "List",
		},
	)
}

// delete User
func (repository *UserRepo) DeleteUser(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteUser(repository.Db, &user, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// now get the User list
	var users []models.User

	err1 := models.GetUsers(repository.Db, &users)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err1})
		return
	}
	c.HTML(http.StatusOK,
		"user_list.html",
		gin.H{
			"users":    users,
			"subtitle": "List",
		},
	)
}

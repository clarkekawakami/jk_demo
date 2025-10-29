package main

import (
	"jk_demo/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":3000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")
	r.StaticFile("/favicon.ico", "./public/favicon.ico")

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,
			"index.html",
			gin.H{
				"title":  "Home Page",
				"active": "home",
			},
		)
	})

	// r.GET("/about", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK,
	// 		"about_page.html",
	// 		gin.H{
	// 			"title":  "About Me",
	// 			"active": "about",
	// 		},
	// 	)
	// })

	userRepo := controllers.New()
	// Users
	r.GET("/users", userRepo.GetUsersPage)
	// xhr endpoints to get htmx html fragments
	r.POST("/users/0", userRepo.CreateUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)
	// r.GET("/userslist", userRepo.GetUsersList) // need this?
	r.GET("/user_form/:id", userRepo.GetUserForm)
	// // Facilities
	r.GET("/facilities", userRepo.GetFacilitysPage)
	r.GET("/facility_form/:id", userRepo.GetFacilityForm)
	r.POST("/facilities/0", userRepo.CreateFacility)
	r.DELETE("/facilities/:id", userRepo.DeleteFacility)
	r.PUT("/facilities/:id", userRepo.UpdateFacility)
	// // resources
	r.GET("/resources", userRepo.GetResourcesPage)
	r.GET("/resource_form/:id", userRepo.GetResourceForm)
	r.POST("/resources/0", userRepo.CreateResource)
	r.DELETE("/resources/:id", userRepo.DeleteResource)
	r.PUT("/resources/:id", userRepo.UpdateResource)
	// // appointments
	r.GET("/appointments", userRepo.GetAppointmentsPage)
	r.GET("/appointment_form/:id", userRepo.GetAppointmentForm)
	r.POST("/appointments/0", userRepo.CreateAppointment)
	r.DELETE("/appointments/:id", userRepo.DeleteAppointment)
	r.PUT("/appointments/:id", userRepo.UpdateAppointment)

	return r
}

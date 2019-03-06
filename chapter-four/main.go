package main

import (
	customerController "basic-golang/chapter-four/controllers/customer"
	staffController "basic-golang/chapter-four/controllers/staff"
	screenerFilm "basic-golang/chapter-four/screener/film"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// authenticated group
	auth := r.Group("", gin.BasicAuth(gin.Accounts{"admin": "password"}))
	auth.GET("/customers", customerController.GetCustomers)
	auth.GET("/staffs", staffController.GetStaffs)

	// unauthenticated group
	r.GET("/film/rental", screenerFilm.ScreenFilmRent)

	r.Run(":8483")
}

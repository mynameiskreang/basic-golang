package main

import (
	customerController "basic-golang/chapter-four/controllers/customer"
	staffController "basic-golang/chapter-four/controllers/staff"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	auth := r.Group("", gin.BasicAuth(gin.Accounts{"admin": "password"}))

	auth.GET("/customers", customerController.GetCustomers)
	auth.GET("/staffs", staffController.GetStaffs)
	auth.GET("/staff", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	r.Run(":4200")
}

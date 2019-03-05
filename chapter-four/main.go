package main

import (
	customerController "basic-golang/chapter-four/controllers/customer"
	staffController "basic-golang/chapter-four/controllers/staff"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/customers", customerController.GetCustomers)
	r.GET("/staffs", staffController.GetStaffs)
	r.Run(":4200")
}

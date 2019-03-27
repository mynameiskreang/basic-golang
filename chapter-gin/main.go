package main

import (
	"basic-golang/chapter-gin/app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth := r.Group("/", gin.BasicAuth(gin.Accounts{"admin": "pass"}))

	// Staff
	auth.GET("/staffs", controller.GetStaffs)
	auth.GET("/staff/:id", controller.GetStaff)
	// Customer
	r.GET("/customer", controller.GetCustomersWithQueryParam)
	r.GET("/customer/:id", controller.GetCustomerByID)

	r.GET("/customers/special/concurrency", controller.ProcessCustomers)
	r.GET("/customers/special/sequence", controller.ProcessCustomersSequence)

	r.Run(":8080")

}

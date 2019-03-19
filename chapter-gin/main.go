package main

import (
	"basic-golang/chapter-gin/app/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/customer", controller.GetCustomersWithQueryParam)
	r.GET("/customer/:id", controller.GetCustomerByID)
	r.GET("/staffs", controller.GetStaffs)
	r.GET("/staff/:id", controller.GetStaff)

	r.GET("/customers/special", controller.ProcessCustomers)

	r.Run(":8080")

}

package controller

import (
	"basic-golang/chapter-gin/app/helper"
	"basic-golang/chapter-gin/app/middleware"
	"basic-golang/chapter-gin/app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit := helper.StringToInt(limitStr)

	customers, err := model.GetCustomers(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": len(customers),
		"data":  customers,
	})
}

func GetCustomersWithQueryParam(c *gin.Context) {
	queryParam := middleware.GetQueryParam(c)

	customers, err := model.GetCustomers(queryParam.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": len(customers),
		"data":  customers,
	})
}

func GetCustomerByID(c *gin.Context) {
	param, err := middleware.GetParam(c)
	if err != nil {
		return
	}

	customer, err := model.GetCustomer(param.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}

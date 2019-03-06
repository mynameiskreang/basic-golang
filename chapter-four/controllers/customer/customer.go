package customer

import (
	"basic-golang/chapter-four/models/customer"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCustomers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	customers, err := customer.GetCustomers(limit)
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

package main

import (
	"basic-golang/chapter-four/databases/postgress"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func init() {
	//logrus.SetOutput(os.Stdout)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/customers", func(c *gin.Context) {
		limitStr := c.DefaultQuery("limit", "10")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		customers, err := postgress.GetCustomers(limit)
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
	})
	r.Run(":4200")
}

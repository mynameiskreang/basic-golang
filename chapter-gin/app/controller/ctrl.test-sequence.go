package controller

import (
	"basic-golang/chapter-gin/app/model"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
)

func ProcessCustomersSequence(c *gin.Context) {
	// Start
	startTime := time.Now()

	type Resp struct {
		FullName string
		Value    float64
	}
	resps := []Resp{}
	//type Resps []Resp

	customers, _ := model.GetCustomers(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			result := 0.0
			for _, custo := range customers {
				result += math.Pi * math.Sin(float64(len(custo.FirstName)))
				resp := Resp{custo.FirstName + " " + custo.LastName, math.Abs(result)}
				resps = append(resps, resp)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		result := 0.0
		customer, err := model.GetCustomer(strconv.Itoa(i))
		go func() {
			for j := 0; j < 1000; j++ {
				if err != nil {
					result += math.Pi * math.Sin(float64(len(customer.LastName)))
					resp := Resp{customer.FirstName + " " + customer.LastName, math.Abs(result)}
					resps = append(resps, resp)
				}
			}
		}()
	}

	// Done
	endTime := time.Now()
	diffTIme := endTime.Sub(startTime)

	c.JSON(http.StatusOK, gin.H{
		"time": diffTIme.String(),
	})
}

package controller

import (
	"basic-golang/chapter-gin/app/model"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func ProcessCustomers(c *gin.Context) {
	// Start
	startTime := time.Now()

	type Resp struct {
		FullName string
		Value    float64
	}
	resps := []Resp{}
	//type Resps []Resp

	//groupCustomer := dto.Customers{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		customers, _ := model.GetCustomers(5000)
		for i := 0; i < 1000; i++ {
			result := 0.0
			for _, custo := range customers {
				result += math.Pi * math.Sin(float64(len(custo.FirstName)))
				resp := Resp{custo.FirstName + " " + custo.LastName, math.Abs(result)}
				resps = append(resps, resp)
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			result := 0.0
			for i := 0; i < 1000; i++ {
				customer, _ := model.GetCustomer(strconv.Itoa(i))
				result += math.Pi * math.Sin(float64(len(customer.LastName)))
				//groupCustomer = append(groupCustomer, customer)
				resp := Resp{customer.FirstName + " " + customer.LastName, math.Abs(result)}
				resps = append(resps, resp)
			}
		}
		wg.Done()
	}()

	wg.Wait()

	// Done
	endTime := time.Now()
	diffTIme := endTime.Sub(startTime)

	c.JSON(http.StatusOK, gin.H{
		"time": diffTIme.String(),
	})
}

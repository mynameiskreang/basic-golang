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
	wg.Add(5000)
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
		wg.Done()
	}

	wg.Add(100)

	//ch := make(chan bool, 10)
	for i := 0; i < 100; i++ {
		result := 0.0
		customer, err := model.GetCustomer(strconv.Itoa(i))
		go func() {
			for j := 0; j < 1000; j++ {
				//ch <- true
				if err != nil {
					result += math.Pi * math.Sin(float64(len(customer.LastName)))
					resp := Resp{customer.FirstName + " " + customer.LastName, math.Abs(result)}
					resps = append(resps, resp)
				}
				//<-ch
			}
			wg.Done()
		}()
	}

	wg.Wait()

	// Done
	endTime := time.Now()
	diffTIme := endTime.Sub(startTime)

	c.JSON(http.StatusOK, gin.H{
		"time": diffTIme.String(),
	})
}

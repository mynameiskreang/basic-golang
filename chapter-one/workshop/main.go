package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Product struct {
	ProductID int
	Name      string
	Amount    int
}

// ลอง Uncomment เพื่อดูความแตกต่างของ json response ครับ
/*type Product struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
}*/

var initProduct Product

func init() {
	initProduct.ProductID = rand.Int()
	initProduct.Name = "Init Product"
	initProduct.Amount = 0
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message":"hello 8483"}`))
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		defer log.Printf("method type: %s", method)
		switch method {
		case "GET":
			if loopString := r.URL.Query().Get("loop"); len(loopString) > 0 {
				loopNumber, _ := strconv.Atoi(loopString)
				products := getProducts(loopNumber)
				productsByte, _ := json.Marshal(products)
				w.Write(productsByte)
			} else {
				initProductByte, _ := json.Marshal(initProduct)
				w.Write(initProductByte)
			}
			return
		case "POST":
			w.Write([]byte(`{"message":"Sorry, Currently not supported method POST"}`))
			return
		}

	})

	http.ListenAndServe(":8483", nil)

}

func getProducts(loop int) (products []Product) {
	for i := 0; i < loop; i++ {
		product := Product{}
		product.ProductID = rand.Int()
		product.Name = "product name " + strconv.Itoa(i+1)
		product.Amount = 0
		products = append(products, product)
	}
	return products
}

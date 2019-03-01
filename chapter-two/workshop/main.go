package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Products []Product
type Product struct {
	ProductID int
	Name      string
	Amount    int
}
type ProductsResponse struct {
	Total    int
	Products Products
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
				loopNumber, err := strconv.Atoi(loopString)
				if err != nil {
					w.Write([]byte(`{"message":"loop must be number"}`))
					return
				}
				products := getProducts(loopNumber)

				productsResponse := ProductsResponse{}
				productsResponse.Products = products
				productsResponse.Total = products.Total()

				productsResponseByte, _ := json.Marshal(productsResponse)
				w.Write(productsResponseByte)
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

func getProducts(loop int) (products Products) {
	for i := 0; i < loop; i++ {
		product := Product{}
		product.ProductID = rand.Int()
		product.Name = "product name " + strconv.Itoa(i+1)
		product.Amount = 0
		products = append(products, product)
	}
	return products
}

func (products Products) Total() int {
	return len(products)
}

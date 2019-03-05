package customer

import "time"

type Customers []Customer
type Customer struct {
	CustomerID int       `json:"customer_id"`
	StoreID    int       `json:"store_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	AddressID  int       `json:"address_id"`
	Activebool bool      `json:"activebool"`
	CreateDate time.Time `json:"create_date"`
	LastUpdate time.Time `json:"last_update"`
	Active     int       `json:"active"`
}

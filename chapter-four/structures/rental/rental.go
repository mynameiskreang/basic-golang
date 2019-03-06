package rental

type Rentals []Rental
type Rental struct {
	RentalID    int    `json:"rental_id"`
	RentalDate  string `json:"rental_date"`
	InventoryID int    `json:"inventory_id"`
	CustomerID  int    `json:"customer_id"`
	ReturnDate  string `json:"return_date"`
	StaffID     int    `json:"staff_id"`
	LastUpdate  string `json:"last_update"`
}

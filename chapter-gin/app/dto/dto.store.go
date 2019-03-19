package dto

import "time"

type Stores []Store
type Store struct {
	StoreID        int       `json:"store_id"`
	ManagerStaffID int       `json:"manager_staff_id"`
	AddressID      int       `json:"address_id"`
	LastUpdate     time.Time `json:"last_update"`
}

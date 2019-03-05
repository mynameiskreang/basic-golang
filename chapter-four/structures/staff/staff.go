package staff

type Staffs []Staff
type Staff struct {
	StaffID    int    `json:"staff_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	AddressID  int    `json:"address_id"`
	Email      string `json:"email"`
	StoreID    int    `json:"store_id"`
	Active     bool   `json:"active"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	LastUpdate string `json:"last_update"`
	Picture    string `json:"picture"`
}

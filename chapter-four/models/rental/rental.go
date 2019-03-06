package rental

import (
	"basic-golang/chapter-four/models"
	structRental "basic-golang/chapter-four/structures/rental"
	"github.com/sirupsen/logrus"
)

func GetRentals() (rentals structRental.Rentals, err error) {
	rows, errQuery := models.PgxConn.Query("select * from rental limit 10")
	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "Getrentals",
		}).Error(errQuery)
		return rentals, errQuery
	}
	for rows.Next() {
		rental := structRental.Rental{}

		errScan := rows.Scan(&rental.RentalID, &rental.RentalDate, &rental.InventoryID, &rental.CustomerID, &rental.ReturnDate, &rental.StaffID, &rental.LastUpdate)
		if errScan != nil {
			logrus.WithFields(logrus.Fields{
				"database": "postgress",
				"action":   "Getrentals",
			}).Error(errScan)
		}

		rentals = append(rentals, rental)
	}
	return rentals, rows.Err()
}

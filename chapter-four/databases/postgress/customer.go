package postgress

import (
	structCustomer "basic-golang/chapter-four/structures/customer"
	"github.com/sirupsen/logrus"
)

func GetCustomers(limit int) (customers structCustomer.Customers, err error) {
	rows, errQuery := pgxConn.Query("select * from customer limit $1", limit)
	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "GetCustomers",
		}).Error(errQuery)
		return customers, errQuery
	}
	for rows.Next() {
		customer := structCustomer.Customer{}

		errScan := rows.Scan(&customer.CustomerID, &customer.StoreID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.AddressID, &customer.Activebool, &customer.CreateDate, &customer.LastUpdate, &customer.Active)
		if errScan != nil {
			logrus.WithFields(logrus.Fields{
				"database": "postgress",
				"action":   "GetCustomers",
			}).Error(errScan)
		}

		customers = append(customers, customer)
	}
	return customers, rows.Err()
}

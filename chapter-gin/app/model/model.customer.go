package model

import (
	"basic-golang/chapter-gin/app/dto"
	"github.com/sirupsen/logrus"
)

func GetCustomers(limit int) (customers dto.Customers, err error) {
	rows, errQuery := PgxConn.Query("select * from customer limit $1", limit)
	//defer rows.Close()

	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "GetCustomers",
		}).Error(errQuery)
		return customers, errQuery
	}
	for rows.Next() {
		customer := dto.Customer{}

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

func GetCustomer(customerID string) (customer dto.Customer, err error) {
	rows, errQuery := PgxConn.Query("select * from customer as cus where cus.customer_id=$1", customerID)
	//defer rows.Close()

	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "GetCustomers",
		}).Error(errQuery)
		return customer, errQuery
	}
	for rows.Next() {
		//customer := dto.Customer{}

		errScan := rows.Scan(&customer.CustomerID, &customer.StoreID, &customer.FirstName, &customer.LastName, &customer.Email, &customer.AddressID, &customer.Activebool, &customer.CreateDate, &customer.LastUpdate, &customer.Active)
		if errScan != nil {
			logrus.WithFields(logrus.Fields{
				"database": "postgress",
				"action":   "GetCustomers",
			}).Error(errScan)
		}

		//customers = append(customers, customer)
	}
	return customer, rows.Err()
}

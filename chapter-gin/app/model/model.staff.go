package model

import (
	"basic-golang/chapter-gin/app/dto"
	"fmt"
	"github.com/sirupsen/logrus"
)

func GetStaffs(limit int) (staffs dto.Staffs, err error) {
	rows, errQuery := PgxConn.Query("select staff_id,first_name,last_name from staff limit $1", limit)
	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "GetStaffs",
		}).Error(errQuery)
		return staffs, errQuery
	}
	for rows.Next() {
		staff := dto.Staff{}
		errScan := rows.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName)
		if errScan != nil {
			logrus.WithFields(logrus.Fields{
				"database": "postgress",
				"action":   "GetStaffs",
			}).Error(errScan)
		}

		staffs = append(staffs, staff)
	}
	return staffs, rows.Err()
}

func GetStaff(id string) (staff dto.Staff, err error) {
	rows, errQuery := PgxConn.Query("select staff_id,first_name,last_name from staff where staff_id=$1", id)
	fmt.Println("errQuery", errQuery)
	for rows.Next() {
		errScan := rows.Scan(&staff.StaffID, &staff.FirstName, &staff.LastName)
		if errScan != nil {
			fmt.Println("errScan", errScan)
		}
	}
	return staff, rows.Err()
}

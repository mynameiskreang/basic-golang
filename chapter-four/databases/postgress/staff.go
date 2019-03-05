package postgress

import (
	structuresStaff "basic-golang/chapter-four/structures/staff"
	"github.com/sirupsen/logrus"
)

func GetStaffs(limit int) (staffs structuresStaff.Staffs, err error) {
	rows, errQuery := pgxConn.Query("select staff_id,first_name,last_name from staff limit $1", limit)
	if errQuery != nil {
		defer logrus.WithFields(logrus.Fields{
			"database": "postgress",
			"action":   "GetStaffs",
		}).Error(errQuery)
		return staffs, errQuery
	}
	for rows.Next() {
		staff := structuresStaff.Staff{}
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

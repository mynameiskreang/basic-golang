package staff

import (
	"basic-golang/chapter-four/models/staff"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStaffs(c *gin.Context) {
	staffs, _ := staff.GetStaffs(10)
	c.JSON(http.StatusOK, gin.H{
		"total": len(staffs),
		"data":  staffs,
	})
}

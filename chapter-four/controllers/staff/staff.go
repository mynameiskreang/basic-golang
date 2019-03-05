package staff

import (
	"basic-golang/chapter-four/databases/postgress"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStaffs(c *gin.Context) {
	staffs, _ := postgress.GetStaffs(10)
	c.JSON(http.StatusOK, gin.H{
		"total": len(staffs),
		"data":  staffs,
	})
}

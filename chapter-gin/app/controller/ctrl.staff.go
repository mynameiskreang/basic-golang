package controller

import (
	"basic-golang/chapter-gin/app/helper"
	"basic-golang/chapter-gin/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStaffs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit := helper.StringToInt(limitStr)

	staffs, err := model.GetStaffs(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": len(staffs),
		"data":  staffs,
	})
}

func GetStaff(c *gin.Context) {
	id := c.Param("id")
	staff, _ := model.GetStaff(id)
	// handle staff
	if len(staff.FirstName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no staff " + id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": staff,
	})
}

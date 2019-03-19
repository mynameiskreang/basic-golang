package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Param struct {
	ID string
}

func GetParam(c *gin.Context) (param Param, err error) {
	err = param.getID(c)
	return param, err
}

func (param *Param) getID(c *gin.Context) error {
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be number",
		})
		return err
	}
	param.ID = id
	return err
}

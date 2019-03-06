package film

import (
	filmController "basic-golang/chapter-four/controllers/film"
	"basic-golang/chapter-four/middleware"
	"github.com/gin-gonic/gin"
)

func ScreenFilmRent(c *gin.Context) {
	queryParam := middleware.GetQueryParam(c)
	//spew.Dump(queryParam)
	filmController.GetMoives(c, queryParam)
}

package film

import (
	"basic-golang/chapter-four/middleware"
	"basic-golang/chapter-four/models/film/movie"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMoives(c *gin.Context, param middleware.QueryParam) {
	movies, err := movie.GetMovies()
	fmt.Println(err)
	spew.Dump(movies)
	c.JSON(http.StatusOK, gin.H{
		"total":        len(movies),
		"data":         movies,
		"queryParams ": param,
	})
}

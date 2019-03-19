package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type QueryParam struct {
	Search string
	Filter string
	SortBy string
	Offset int
	Limit  int
}

func GetQueryParam(c *gin.Context) (queryParam QueryParam) {
	queryParam.getSearch(c)
	queryParam.getFilter(c)
	queryParam.getSortBy(c)
	queryParam.getOffset(c)
	queryParam.getLimit(c)
	return queryParam
}

func (queryParam *QueryParam) getSearch(c *gin.Context) {
	if search, isExist := c.GetQuery("search"); isExist {
		queryParam.Search = search
	}
}

func (queryParam *QueryParam) getFilter(c *gin.Context) {
	if filter, isExist := c.GetQuery("filter"); isExist {
		queryParam.Filter = filter
	}
}

func (queryParam *QueryParam) getSortBy(c *gin.Context) {
	if sortBy, isExist := c.GetQuery("sortBy"); isExist {
		queryParam.SortBy = sortBy
	}
}

func (queryParam *QueryParam) getOffset(c *gin.Context) {
	if offset, isExist := c.GetQuery("offset"); isExist {
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "offset must be integer",
			})
			return
		}
		queryParam.Offset = offsetInt
	}
}

func (queryParam *QueryParam) getLimit(c *gin.Context) {
	if offsetStr, isExist := c.GetQuery("limit"); isExist {

		limitInt, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "limit must be integer",
			})
			return
		}

		queryParam.Limit = limitInt
	}
}

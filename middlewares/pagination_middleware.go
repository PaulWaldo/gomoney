package middlewares

import (
	// "clean-architecture/constants"
	// "clean-architecture/lib"
	"log"
	"strconv"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	// logger lib.Logger
	logger log.Logger
}

func NewPaginationMiddleware(logger log.Logger) PaginationMiddleware {
	return PaginationMiddleware{logger: logger}
}

func /*(p PaginationMiddleware)*/ Paginator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// p.logger.Output("setting up pagination middleware")

		perPage, err := strconv.ParseInt(c.Query("per_page"), 10, 0)
		if err != nil {
			perPage = 10
		}

		page, err := strconv.ParseInt(c.Query("page"), 10, 0)
		if err != nil {
			page = 0
		}

		c.Set(constants.Limit, perPage)
		c.Set(constants.Page, page)

		c.Next()
	}
}

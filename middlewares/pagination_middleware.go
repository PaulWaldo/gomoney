package middlewares

import (
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
		// fmt.Println("********* Query Params *************")
		// for k, v := range c.Request.URL.Query() {
		// 	fmt.Printf("%s = %s\n", k, v)
		// }

		// Designed to handle parameters as sent by DataTables
		// https://datatables.net/manual/server-side#Sent-parameters
		perPage, err := strconv.ParseInt(c.Query("length"), 10, 0)
		if err != nil {
			perPage = 10
		}

		offset, err := strconv.ParseInt(c.Query("start"), 10, 0)
		if err != nil {
			offset = 0
		}

		c.Set(constants.Limit, perPage)
		c.Set(constants.Offset, offset)

		c.Next()
	}
}

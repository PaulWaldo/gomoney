package routes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/PaulWaldo/gomoney/middlewares"
	"github.com/PaulWaldo/gomoney/utils"
	"github.com/gin-gonic/gin"
)

// func jSONWithPagination(c *gin.Context, statusCode int, response utils.PaginatedResponse) {
// 	limit, _ := c.MustGet(constants.Limit).(int64)
// 	offset, _ := c.MustGet(constants.Offset).(int64)
// 	response.HasNext = (response.Count - limit*offset) > 0

// 	c.JSON(statusCode, response)
// }

func (controller Controller) transactionsListHandler(c *gin.Context) {
	controller.services.Transaction.SetPaginationScope(utils.Paginate(c))
	txns, count, err := controller.services.Transaction.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	response := utils.PaginatedResponse{Data: txns, Count: count}
	// jSONWithPagination(c, http.StatusOK, response)
	limit, _ := c.MustGet(constants.Limit).(int64)
	offset, _ := c.MustGet(constants.Offset).(int64)
	response.HasNext = (response.Count - limit*offset) > 0

	drawStr := c.Query("draw")
	draw, err := strconv.Atoi(drawStr)
	if err != nil {
		log.Fatalf("draw parameter from frontend '%s' unable to be converted to int: %s", drawStr, err)
		c.JSON(http.StatusInternalServerError, "error")
	}
	// https://datatables.net/manual/server-side#Returned-data
	c.JSON(http.StatusOK, gin.H{
		"draw":            draw,
		"recordsTotal":    response.Count,
		"data":            response.Data,
		"recordsFiltered": response.Count,
	})
}

func (controller Controller) AddTransactionRoutes() {
	group := controller.router.Group(constants.TransactionsURL).Use(middlewares.Paginator())
	{
		group.GET("", controller.transactionsListHandler)
	}
}

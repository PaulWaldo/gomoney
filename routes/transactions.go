package routes

import (
	"net/http"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/PaulWaldo/gomoney/middlewares"
	"github.com/PaulWaldo/gomoney/utils"
	"github.com/gin-gonic/gin"
)

func jSONWithPagination(c *gin.Context, statusCode int, response utils.PaginatedResponse) {
	limit, _ := c.MustGet(constants.Limit).(int64)
	size, _ := c.MustGet(constants.Page).(int64)

	c.JSON(
		statusCode,
		gin.H{
			"data": response.Data,
			"pagination": gin.H{
				"has_next": (response.Count - limit*size) > 0,
				"count":    response.Count,
			},
		})
}

func (controller Controller) transactionsListHandler(c *gin.Context) {
	controller.services.Transaction.SetPaginationScope(utils.Paginate(c))
	response, err := controller.services.Transaction.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	jSONWithPagination(c, http.StatusOK, response)
}

func (controller Controller) AddTransactionRoutes() {
	group := controller.router.Group(constants.TransactionsURL).Use(middlewares.Paginator())
	{
		group.GET("", controller.transactionsListHandler)
	}
}

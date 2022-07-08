package routes

import (
	"net/http"

	"github.com/PaulWaldo/gomoney/constants"
	"github.com/PaulWaldo/gomoney/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller Controller) transactionsListHandler(c *gin.Context) {
	transactions, err := controller.services.Transaction.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, transactions)
}

func (controller Controller) AddTransactionRoutes() {
	group := controller.router.Group(constants.TransactionsURL).Use(middlewares.Paginator())
	{
		group.GET("", controller.transactionsListHandler)
	}
}

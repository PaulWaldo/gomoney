package routes

import (
	"net/http"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowHandler(c *gin.Context) {
	// accounts = dom
	c.HTML(http.StatusOK, "cashflow", gin.H{
		"PageTitle": "MoneyMinder - Cashflow",
		"Accounts": []domain.Account {domain.NewAccount("fred", domain.Checking)},
	})
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET("/cashflow", controller.cashFlowHandler)
}

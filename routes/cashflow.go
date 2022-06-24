package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowHandler(c *gin.Context) {
	var status = http.StatusOK
	accounts, err := controller.services.Account.List()
	if err != nil {
		status = http.StatusInternalServerError
	}
	err = errors.New("Yikes!")
	c.HTML(status, "cashflow", gin.H{
		"PageTitle": "MoneyMinder - Cashflow",
		"Error":     err,
		"Accounts":  accounts,
	})
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET("/cashflow", controller.cashFlowHandler)
}

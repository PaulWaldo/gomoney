package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowHandler(c *gin.Context) {
	var status = http.StatusOK
	accounts, err := controller.services.Account.List()
	if err != nil {
		status = http.StatusInternalServerError
	}
	// err = errors.New("Yikes!")
	c.HTML(status, "base.html", gin.H{
		"PageTitle": "MoneyMinder - Cashflow",
		"Error":     err,
		"Accounts":  accounts,
	})
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET("/cashflow", controller.cashFlowHandler)
	controller.router.GET("/", controller.cashFlowHandler)
	controller.router.GET("/index.html", controller.cashFlowHandler)
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const cashflowURL = "/cashflow"

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

func (controller Controller) cashFlowRedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, cashflowURL)
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET(cashflowURL, controller.cashFlowHandler)
	controller.router.GET("/", controller.cashFlowRedirectHandler)
	controller.router.GET("/index.html", controller.cashFlowRedirectHandler)
}

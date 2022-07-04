package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowAllAccountsHandler(c *gin.Context) {
	var status = http.StatusOK
	accounts, err := controller.services.Account.List()
	if err != nil {
		status = http.StatusInternalServerError
	}
	c.HTML(status, "base.html", gin.H{
		"PageTitle":         "MoneyMinder - Cashflow",
		"Error":             err,
		"Accounts":          accounts,
		"SelectedAccountID": 0,
	})
}

func (controller Controller) cashFlowSpecificAccountsHandler(c *gin.Context) {
	var status = http.StatusOK
	accountIdParam := c.Param("accountId")
	accountId, err := strconv.ParseUint(accountIdParam, 10, 32)
	if err != nil {
		c.HTML(status, "base.html", gin.H{
			"PageTitle":         "MoneyMinder - Cashflow",
			"Error":             err,
			"Accounts":          []models.Account{},
			"SelectedAccountID": accountId,
		})
	}
	fmt.Println(accountId)
	account, err := controller.services.Account.Get(uint(accountId))
	if err != nil {
		status = http.StatusInternalServerError
	}
	c.HTML(status, "base.html", gin.H{
		"PageTitle":         "MoneyMinder - Cashflow",
		"Error":             err,
		"Accounts":          []models.Account{account},
		"SelectedAccountID": accountId,
	})
}

func (controller Controller) cashFlowRedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/cashflow")
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.LoadHTMLGlob("templates/*")
	controller.router.GET("/cashflow", controller.cashFlowAllAccountsHandler)
	controller.router.GET("/cashflow/:accountId", controller.cashFlowSpecificAccountsHandler)
	controller.router.GET("/", controller.cashFlowRedirectHandler)
	controller.router.GET("/index.html", controller.cashFlowRedirectHandler)
}

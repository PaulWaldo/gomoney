package routes

import (
	"net/http"
	"strconv"

	"github.com/PaulWaldo/gomoney/pkg/domain/models"
	"github.com/gin-gonic/gin"
)

func (controller Controller) returnError(c *gin.Context, err error) {
	c.HTML(http.StatusInternalServerError, "base.html", gin.H{
		"PageTitle":         "MoneyMinder - Cashflow",
		"Error":             err,
		"Accounts":          []models.Account{},
		"SelectedAccountID": 0,
	})
}

func (controller Controller) cashFlowAllAccountsHandler(c *gin.Context) {
	var status = http.StatusOK
	accountIdParam := c.Param("accountId")
	var accountId = uint64(0)
	var err error
	if len(accountIdParam) > 0 {
		accountId, err = strconv.ParseUint(accountIdParam, 10, 32)
		if err != nil {
			controller.returnError(c, err)
		}
	}

	accounts, err := controller.services.Account.List()
	if err != nil {
		controller.returnError(c, err)
	}

	c.HTML(status, "base.html", gin.H{
		"PageTitle":         "MoneyMinder - Cashflow",
		"Error":             nil,
		"Accounts":          accounts,
		"SelectedAccountID": accountId,
	})
}

// func (controller Controller) cashFlowSpecificAccountsHandler(c *gin.Context) {
// 	var status = http.StatusOK
// 	accountIdParam := c.Param("accountId")
// 	accountId, err := strconv.ParseUint(accountIdParam, 10, 32)
// 	if err != nil {
// 		c.HTML(status, "base.html", gin.H{
// 			"PageTitle":         "MoneyMinder - Cashflow",
// 			"Error":             err,
// 			"Accounts":          []models.Account{},
// 			"SelectedAccountID": accountId,
// 		})
// 	}
// 	fmt.Println(accountId)
// 	account, err := controller.services.Account.Get(uint(accountId))
// 	if err != nil {
// 		status = http.StatusInternalServerError
// 	}
// 	c.HTML(status, "base.html", gin.H{
// 		"PageTitle":         "MoneyMinder - Cashflow",
// 		"Error":             err,
// 		"Accounts":          []models.Account{account},
// 		"SelectedAccountID": accountId,
// 	})
// }

func (controller Controller) cashFlowRedirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/cashflow")
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.LoadHTMLGlob("templates/*")
	group := controller.router.Group("/cashflow")
	{
		group.GET("", controller.cashFlowAllAccountsHandler)
		group.GET("/:accountId", controller.cashFlowAllAccountsHandler)
	}
	controller.router.GET("/", controller.cashFlowRedirectHandler)
	controller.router.GET("/index.html", controller.cashFlowRedirectHandler)
}

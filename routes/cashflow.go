package routes

import (
	"net/http"

	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowHandler(c *gin.Context) {
	// controller.
	c.HTML(http.StatusOK, "cashflow.gohtml", gin.H{
		"PageTitle": "Main website",
		"Accounts": []domain.Account {domain.NewAccount("fred", domain.Checking)},
	})
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET("/cashflow", controller.cashFlowHandler)
}

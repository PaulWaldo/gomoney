package routes

import (
	"net/http"

	"github.com/PaulWaldo/gomoney/internal/db/models"
	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "cashflow", gin.H{
		"PageTitle": "MoneyMinder - Cashflow",
		"Accounts": []models.Account{
			{Name: "fred", Type: models.Checking},
		},
	})
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET("/cashflow", controller.cashFlowHandler)
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) cashFlowHandler(c *gin.Context) {
	// controller.
	c.HTML(http.StatusOK, "accounts.gohtml", gin.H{
		"title": "Main website",
	})
}

func (controller Controller) AddCashFlowRoutes() {
	controller.router.GET("/cashflow", controller.cashFlowHandler)
}
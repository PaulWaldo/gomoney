package main

import (
	// "log"

	// "github.com/PaulWaldo/gomoney/internal/db"
	// "github.com/PaulWaldo/gomoney/routes"
	"github.com/PaulWaldo/gomoney/ui"
	// "github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

func main() {
	ui.RunApp()
	// gin.SetMode(gin.DebugMode)
	// r := gin.Default()

	// r.Static("/myjs", "js")
	// r.Static("/static", "node_modules/startbootstrap-sb-admin-2")
	// services, _, err := db.NewSqliteInMemoryServices(&gorm.Config{}, true)
	// if err != nil {
	// 	panic(err)
	// }

	// controller := routes.NewController(r, services)
	// controller.AddCashFlowRoutes()
	// controller.AddTransactionRoutes()

	// log.Print("Starting server on port 8080")
	// r.Run(":8080")
}

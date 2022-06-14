package main

import (
	"fmt"
	"log"

	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/routes"
	"github.com/gin-gonic/gin"
)

// type Routable interface {
// 	AddRoutes(r *gin.Engine)
// }

func main() {
	db, err := db.ConnectToDatabase()
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %s", err))
	}

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.LoadHTMLGlob("../../internal/html/*")
	controller := routes.NewController(db, r)

	controller.AddCashFlowRoutes()

	log.Print("Starting server")
	r.Run(":8080")
}

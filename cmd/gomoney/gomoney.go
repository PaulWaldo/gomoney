package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PaulWaldo/gomoney/internal/db"
	"github.com/PaulWaldo/gomoney/routes"
	"github.com/gin-gonic/gin"
)

// type Routable interface {
// 	AddRoutes(r *gin.Engine)
// }

func main() {
	// db, err := db.ConnectToDatabase()
	// if err != nil {
	// 	panic(fmt.Sprintf("Unable to connect to database: %s", err))
	// }

	// gin.SetMode(gin.DebugMode)
	r := gin.Default()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("CWD is %s", cwd)
	r.LoadHTMLGlob("../../templates/*")
	r.Static("/static", "../../node_modules/startbootstrap-sb-admin-2")
	services, err := db.NewSqliteInMemoryServices()
	if err != nil {
		panic(err)
	}

	controller := routes.NewController(r, services)

	controller.AddCashFlowRoutes()

	log.Print("Starting server on port 8080")
	r.Run(":8080")
}

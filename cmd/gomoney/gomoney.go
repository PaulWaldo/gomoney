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
	db, err := db.ConnectToDatabase()
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %s", err))
	}

	// gin.SetMode(gin.DebugMode)
	r := gin.Default()

	cwd, err:=os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("CWD is %s", cwd)

	r.LoadHTMLGlob("../../internal/templates/*.gohtml")
	controller := routes.NewController(db, r)

	controller.AddCashFlowRoutes()

	log.Print("Starting server")
	r.Run(":8080")
}

// Package http handles all routes
package routes

import (
	"github.com/PaulWaldo/gomoney/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	// db       *gorm.DB
	router   *gin.Engine
	services *service.Services
}

func NewController(router *gin.Engine, services *service.Services) Controller {
	return Controller{
		// db:       services.,
		router:   router,
		services: services,
	}
}

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	return r
// }

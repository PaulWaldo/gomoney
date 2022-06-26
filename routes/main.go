// Package http handles all routes
package routes

import (
	"github.com/PaulWaldo/gomoney/pkg/domain"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	// db       *gorm.DB
	router   *gin.Engine
	services *domain.Services
}

func NewController(router *gin.Engine, services *domain.Services) Controller {
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

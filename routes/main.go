// Package http handles all routes
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	db     *gorm.DB
	router *gin.Engine
}

func NewController(db *gorm.DB, router *gin.Engine) Controller {
	return Controller{
		db: db, router: router,
	}
}

// func NewRouter() *gin.Engine {
// 	r := gin.Default()
// 	return r
// }

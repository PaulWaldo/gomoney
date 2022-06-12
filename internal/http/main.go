// Package http handles all routes
package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		db: db,
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	return r
}
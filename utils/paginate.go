package utils

import (
	"github.com/PaulWaldo/gomoney/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit, _ := c.MustGet(constants.Limit).(int64)
		offset, _ := c.MustGet(constants.Offset).(int64)
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

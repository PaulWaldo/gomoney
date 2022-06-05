package middleware

import (
	"log"

	"github.com/PaulWaldo/gomoney/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Session(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionIdentifierInterface := session.Get(SessionIdentifierKey)

		if sessionIdentifier, ok := sessionIdentifierInterface.(string); ok {
			ses := models.Session{
				Identifier: sessionIdentifier,
			}
			res := db.Where(&ses).First(&ses)
			if res.Error == nil {
				c.Set(UserIDKey, ses.UserID)
			} else {
				log.Println(res.Error)
			}
		}
		c.Next()
	}
}

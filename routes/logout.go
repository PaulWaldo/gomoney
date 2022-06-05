package routes

import (
	"log"
	"net/http"

	"github.com/PaulWaldo/gomoney/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (controller Controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(middleware.SessionIdentifierKey)
	err := session.Save()
	log.Println(err)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

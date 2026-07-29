package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowHome(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("user_email")
	name := session.Get("user_name")

	if email == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"User": gin.H{
			"Name":  name,
			"Email": email,
		},
	})
}
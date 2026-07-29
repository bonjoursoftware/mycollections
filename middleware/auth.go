package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// RequireAuth ensures the user is logged in before accessing a route.
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userEmail := session.Get("user_email")
		if userEmail == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("user_email", userEmail)
		c.Next()
	}
}
package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("unauthenticated", func(t *testing.T) {
		r := gin.Default()
		store := cookie.NewStore([]byte("secret"))
		r.Use(sessions.Sessions("mysession", store))
		r.Use(RequireAuth())
		r.GET("/protected", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})

		req, _ := http.NewRequest("GET", "/protected", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		if w.Code != http.StatusFound {
			t.Errorf("Expected status %d, got %d", http.StatusFound, w.Code)
		}
		if w.Header().Get("Location") != "/login" {
			t.Errorf("Expected redirection to /login, got %s", w.Header().Get("Location"))
		}
	})

	t.Run("authenticated", func(t *testing.T) {
		r := gin.Default()
		store := cookie.NewStore([]byte("secret"))
		r.Use(sessions.Sessions("mysession", store))

		// Login route to set session
		r.GET("/login-mock", func(c *gin.Context) {
			session := sessions.Default(c)
			session.Set("user_email", "test@example.com")
			session.Save()
			c.String(http.StatusOK, "Logged in")
		})

		r.Use(RequireAuth())
		r.GET("/protected", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})

		// 1. Mock Login
		req, _ := http.NewRequest("GET", "/login-mock", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		
		cookie := w.Result().Header.Get("Set-Cookie")

		// 2. Access protected
		req2, _ := http.NewRequest("GET", "/protected", nil)
		req2.Header.Add("Cookie", cookie)
		w2 := httptest.NewRecorder()
		r.ServeHTTP(w2, req2)

		if w2.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w2.Code)
		}
	})
}
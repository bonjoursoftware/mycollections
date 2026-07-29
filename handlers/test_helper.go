package handlers

import (
	"mycollections/db"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func setupTestRouter(t *testing.T) *gin.Engine {
	// Use in-memory SQLite DB
	err := db.InitDB("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to init db: %v", err)
	}

	// Clean up db
	_, _ = db.DB.Exec("DELETE FROM users")

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Initialize Session Store
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Find the templates directory from the project root
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	templatesPath := filepath.Join(basepath, "..", "templates", "*")
	
	// Load HTML templates
	r.LoadHTMLGlob(templatesPath)

	return r
}
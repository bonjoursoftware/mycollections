package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"mycollections/db"
	"mycollections/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Helper to simulate a logged-in user session
func mockAuthMiddleware(email string, userID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user_email", email)
		session.Set("user_id", userID)
		_ = session.Save()
		c.Next()
	}
}

func TestShowCollections_Unauthenticated(t *testing.T) {
	router := setupTestRouter(t)
	// Notice we don't add the mock auth middleware, so session is empty
	router.GET("/collections", ShowCollections)

	req, _ := http.NewRequest(http.MethodGet, "/collections", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}
	if w.Header().Get("Location") != "/login" {
		t.Fatalf("Expected redirect to /login, got %s", w.Header().Get("Location"))
	}
}

func TestShowCollections_Authenticated(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")
	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "My Collection")

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.GET("/collections", ShowCollections)

	req, _ := http.NewRequest(http.MethodGet, "/collections", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "My Collection") {
		t.Errorf("Expected body to contain 'My Collection', but it did not. Body: %s", body)
	}
}

func TestHandleCreateCollection_Unauthenticated(t *testing.T) {
	router := setupTestRouter(t)
	router.POST("/collections", HandleCreateCollection)

	req, _ := http.NewRequest(http.MethodPost, "/collections", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}
}

func TestHandleCreateCollection_Success(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")
	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections", HandleCreateCollection)

	form := url.Values{}
	form.Add("name", "New Collection")
	req, _ := http.NewRequest(http.MethodPost, "/collections", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}
	if w.Header().Get("Location") != "/collections" {
		t.Fatalf("Expected redirect to /collections, got %s", w.Header().Get("Location"))
	}

	// Verify it was created
	collections, _ := models.GetCollections(u.ID)
	if len(collections) != 1 || collections[0].Name != "New Collection" {
		t.Fatalf("Expected 1 collection named 'New Collection', got: %+v", collections)
	}
}

func TestHandleCreateCollection_EmptyName(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")
	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections", HandleCreateCollection)

	form := url.Values{}
	form.Add("name", "") // Empty name
	req, _ := http.NewRequest(http.MethodPost, "/collections", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status 400 Bad Request, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Collection name cannot be empty") {
		t.Errorf("Expected body to contain error message about empty name. Body: %s", body)
	}
}

func TestHandleCreateCollection_DuplicateName(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")
	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "Existing Collection")

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections", HandleCreateCollection)

	form := url.Values{}
	form.Add("name", "Existing Collection")
	req, _ := http.NewRequest(http.MethodPost, "/collections", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status 400 Bad Request, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "A collection with that name already exists") {
		t.Errorf("Expected body to contain duplicate error message. Body: %s", body)
	}
}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"mycollections/db"
	"mycollections/models"
)

func TestShowCollectionTags_Success(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "Test Collection")
	cols, _ := models.GetCollections(u.ID)
	c := cols[0]
	_ = models.CreateTag(c.ID, "Sci-Fi")

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.GET("/collections/:id", ShowCollectionTags)

	req, _ := http.NewRequest(http.MethodGet, "/collections/"+strconv.Itoa(c.ID), nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Sci-Fi") {
		t.Errorf("Expected body to contain tag 'Sci-Fi'")
	}
}

func TestHandleCreateTag_Success(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "Test Collection")
	cols, _ := models.GetCollections(u.ID)
	c := cols[0]

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections/:id/tags", HandleCreateTag)

	form := url.Values{}
	form.Add("name", "New-Tag")
	req, _ := http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}

	tags, _ := models.GetTags(c.ID)
	if len(tags) != 1 || tags[0].Name != "New-Tag" {
		t.Fatalf("Expected tag to be created")
	}
}

func TestHandleCreateTag_Invalid(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "Test Collection")
	cols, _ := models.GetCollections(u.ID)
	c := cols[0]

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections/:id/tags", HandleCreateTag)

	form := url.Values{}
	form.Add("name", "Invalid Tag Space")
	req, _ := http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status 400 Bad Request, got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "alphanumeric") {
		t.Errorf("Expected body to contain invalid tag name error")
	}
}

func TestHandleEditTag_Success(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "Test Collection")
	cols, _ := models.GetCollections(u.ID)
	c := cols[0]
	_ = models.CreateTag(c.ID, "Old-Tag")
	
	tags, _ := models.GetTags(c.ID)
	tagID := tags[0].ID

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections/:id/tags/:tag_id/edit", HandleEditTag)

	form := url.Values{}
	form.Add("name", "Updated-Tag")
	req, _ := http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags/"+strconv.Itoa(tagID)+"/edit", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}

	updatedTags, _ := models.GetTags(c.ID)
	if updatedTags[0].Name != "Updated-Tag" {
		t.Fatalf("Expected tag to be updated")
	}
}

func TestHandleDeleteTag_Success(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = models.CreateUser("test@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("test@example.com", "password")
	_ = models.CreateCollection(u.ID, "Test Collection")
	cols, _ := models.GetCollections(u.ID)
	c := cols[0]
	_ = models.CreateTag(c.ID, "Tag-To-Delete")
	
	tags, _ := models.GetTags(c.ID)
	tagID := tags[0].ID

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	router.POST("/collections/:id/tags/:tag_id/delete", HandleDeleteTag)

	req, _ := http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags/"+strconv.Itoa(tagID)+"/delete", nil)
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}

	remainingTags, _ := models.GetTags(c.ID)
	if len(remainingTags) != 0 {
		t.Fatalf("Expected tag to be deleted")
	}
}

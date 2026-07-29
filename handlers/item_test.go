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

func TestItemHandlers(t *testing.T) {
	router := setupTestRouter(t)
	_, _ = db.DB.Exec("DELETE FROM item_tags")
	_, _ = db.DB.Exec("DELETE FROM items")
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = models.CreateUser("itemuser@example.com", "Test User", "password")
	u, _ := models.AuthenticateUser("itemuser@example.com", "password")
	_ = models.CreateCollection(u.ID, "Test Collection")
	cols, _ := models.GetCollections(u.ID)
	c := cols[0]
	_ = models.CreateTag(c.ID, "TagA")
	_ = models.CreateTag(c.ID, "TagB")
	tags, _ := models.GetTags(c.ID)
	tagA, tagB := tags[0], tags[1]

	router.Use(mockAuthMiddleware(u.Email, u.ID))
	
	router.GET("/collections/:id/tags/:tag_id", ShowTagItems)
	router.POST("/collections/:id/tags/:tag_id/items", HandleCreateItem)
	router.POST("/collections/:id/tags/:tag_id/items/:item_id/edit", HandleEditItem)
	router.POST("/collections/:id/tags/:tag_id/items/:item_id/delete", HandleDeleteItem)

	// Test ShowTagItems (empty)
	req, _ := http.NewRequest(http.MethodGet, "/collections/"+strconv.Itoa(c.ID)+"/tags/"+strconv.Itoa(tagA.ID), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", w.Code)
	}

	// Test HandleCreateItem
	form := url.Values{}
	form.Add("name", "New Item")
	form.Add("description", "A description")
	form.Add("tags", strconv.Itoa(tagA.ID))
	form.Add("tags", strconv.Itoa(tagB.ID))
	
	req, _ = http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags/"+strconv.Itoa(tagA.ID)+"/items", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}

	items, _ := models.GetItemsByTag(tagA.ID)
	if len(items) != 1 || items[0].Name != "New Item" {
		t.Fatalf("Expected 1 item named 'New Item'")
	}
	itemID := items[0].ID

	// Test HandleEditItem
	editForm := url.Values{}
	editForm.Add("name", "Edited Item")
	editForm.Add("description", "Edited desc")
	editForm.Add("tags", strconv.Itoa(tagB.ID)) // Removing tagA

	req, _ = http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags/"+strconv.Itoa(tagA.ID)+"/items/"+strconv.Itoa(itemID)+"/edit", strings.NewReader(editForm.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}

	itemsAfterEdit, _ := models.GetItemsByTag(tagA.ID)
	if len(itemsAfterEdit) != 0 {
		t.Fatalf("Expected item to be removed from tagA")
	}

	itemsInTagB, _ := models.GetItemsByTag(tagB.ID)
	if len(itemsInTagB) != 1 || itemsInTagB[0].Name != "Edited Item" {
		t.Fatalf("Expected edited item to be in tagB")
	}

	// Test HandleDeleteItem
	req, _ = http.NewRequest(http.MethodPost, "/collections/"+strconv.Itoa(c.ID)+"/tags/"+strconv.Itoa(tagA.ID)+"/items/"+strconv.Itoa(itemID)+"/delete", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("Expected status 302 Found, got %d", w.Code)
	}

	itemsFinal, _ := models.GetItemsByTag(tagB.ID)
	if len(itemsFinal) != 0 {
		t.Fatalf("Expected item to be deleted globally")
	}
}

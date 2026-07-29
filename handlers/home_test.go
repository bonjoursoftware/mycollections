package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
    
	"mycollections/middleware"
)

func TestShowHome_Unauthenticated(t *testing.T) {
	r := setupTestRouter(t)
	r.GET("/", ShowHome)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Since there's no auth middleware directly on this route in this test setup,
	// the handler will redirect if session email is nil.
	if w.Code != http.StatusFound {
		t.Errorf("Expected status %d, got %d", http.StatusFound, w.Code)
	}
	if w.Header().Get("Location") != "/login" {
		t.Errorf("Expected redirection to /login, got %s", w.Header().Get("Location"))
	}
}

func TestShowHome_Authenticated(t *testing.T) {
	r := setupTestRouter(t)
	
	// Create user
	r.POST("/register", HandleRegister)
	r.POST("/login", HandleLogin)
	
	authorized := r.Group("/")
	authorized.Use(middleware.RequireAuth())
	{
		authorized.GET("/", ShowHome)
	}

	// 1. Register
	form := url.Values{}
	form.Add("name", "Home User")
	form.Add("email", "home@example.com")
	form.Add("password", "secret")
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 2. Login
	loginForm := url.Values{}
	loginForm.Add("email", "home@example.com")
	loginForm.Add("password", "secret")
	req, _ = http.NewRequest("POST", "/login", strings.NewReader(loginForm.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	
	// Extract cookie (Wait, http/httptest multiple Set-Cookie headers need care)
	cookies := w.Result().Cookies()
	var sessionCookie string
	for _, c := range cookies {
		if c.Name == "mysession" {
			sessionCookie = c.String()
			break
		}
	}

	// 3. Access Home
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Add("Cookie", sessionCookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	body := w.Body.String()
	if !strings.Contains(body, "Dashboard") {
		t.Errorf("Expected response to contain 'Dashboard'")
	}
	if !strings.Contains(body, "Home User") {
		t.Errorf("Expected response to contain user's name")
	}
}
package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestShowRegister(t *testing.T) {
	r := setupTestRouter(t)
	r.GET("/register", ShowRegister)

	req, _ := http.NewRequest("GET", "/register", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	if !strings.Contains(w.Body.String(), "Register") {
		t.Errorf("Expected response to contain 'Register'")
	}
}

func TestHandleRegister(t *testing.T) {
	r := setupTestRouter(t)
	r.POST("/register", HandleRegister)

	form := url.Values{}
	form.Add("name", "Test User")
	form.Add("email", "test@example.com")
	form.Add("password", "password123")

	req, _ := http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	if !strings.Contains(w.Body.String(), "Registration successful") {
		t.Errorf("Expected response to contain success message")
	}

	// Test duplicate email
	req, _ = http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !strings.Contains(w.Body.String(), "Email already exists") {
		t.Errorf("Expected response to contain duplicate email error")
	}
}

func TestHandleLogin(t *testing.T) {
	r := setupTestRouter(t)
	r.POST("/register", HandleRegister)
	r.POST("/login", HandleLogin)

	// Register a user first
	form := url.Values{}
	form.Add("name", "Login User")
	form.Add("email", "login@example.com")
	form.Add("password", "secret")

	req, _ := http.NewRequest("POST", "/register", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Test valid login
	loginForm := url.Values{}
	loginForm.Add("email", "login@example.com")
	loginForm.Add("password", "secret")

	req, _ = http.NewRequest("POST", "/login", strings.NewReader(loginForm.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusFound { // Should redirect
		t.Errorf("Expected status %d, got %d", http.StatusFound, w.Code)
	}
	if w.Header().Get("Location") != "/" {
		t.Errorf("Expected redirection to /, got %s", w.Header().Get("Location"))
	}
	cookie := w.Header().Get("Set-Cookie")
	if !strings.Contains(cookie, "mysession") {
		t.Errorf("Expected session cookie to be set")
	}

	// Test invalid login
	loginForm.Set("password", "wrong")
	req, _ = http.NewRequest("POST", "/login", strings.NewReader(loginForm.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
	if !strings.Contains(w.Body.String(), "Invalid email or password") {
		t.Errorf("Expected response to contain invalid credentials error")
	}
}

func TestHandleLogout(t *testing.T) {
	r := setupTestRouter(t)
	r.POST("/logout", HandleLogout)

	req, _ := http.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Errorf("Expected status %d, got %d", http.StatusFound, w.Code)
	}
	if w.Header().Get("Location") != "/login" {
		t.Errorf("Expected redirection to /login, got %s", w.Header().Get("Location"))
	}
}
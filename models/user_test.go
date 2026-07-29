package models

import (
	"testing"

	"mycollections/db"
)

// setupTestDB creates an in-memory SQLite database for testing.
func setupTestDB(t *testing.T) {
	err := db.InitDB("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("Failed to initialize test db: %v", err)
	}

	// Clean up table before each test run if needed (though memory DBs are ephemeral)
	_, err = db.DB.Exec("DELETE FROM users")
	if err != nil {
		t.Fatalf("Failed to clean users table: %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	setupTestDB(t)

	err := CreateUser("test@example.com", "Test User", "password123")
	if err != nil {
		t.Fatalf("Expected nil, got error: %v", err)
	}

	// Test Duplicate
	err = CreateUser("test@example.com", "Test User 2", "password456")
	if err != ErrEmailExists {
		t.Fatalf("Expected ErrEmailExists, got %v", err)
	}
}

func TestAuthenticateUser(t *testing.T) {
	setupTestDB(t)

	email := "auth@example.com"
	name := "Auth User"
	password := "secret"

	err := CreateUser(email, name, password)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Happy path
	user, err := AuthenticateUser(email, password)
	if err != nil {
		t.Fatalf("Failed to authenticate valid user: %v", err)
	}
	if user.Email != email || user.Name != name {
		t.Errorf("Returned user data mismatch: %+v", user)
	}

	// Wrong password
	_, err = AuthenticateUser(email, "wrongpassword")
	if err != ErrInvalidCredentials {
		t.Errorf("Expected ErrInvalidCredentials, got %v", err)
	}

	// Wrong email
	_, err = AuthenticateUser("nonexistent@example.com", password)
	if err != ErrInvalidCredentials {
		t.Errorf("Expected ErrInvalidCredentials, got %v", err)
	}
}

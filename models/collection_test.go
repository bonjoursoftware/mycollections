package models

import (
	"mycollections/db"
	"testing"
)

func TestCreateCollection(t *testing.T) {
	setupTestDB(t)
	// Clear collections table just in case
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")
	_ = CreateUser("test@example.com", "Test User", "password")
	u, _ := AuthenticateUser("test@example.com", "password")

	userID := u.ID
	
	// Happy path
	err := CreateCollection(userID, "My Books")
	if err != nil {
		t.Fatalf("Expected nil, got error: %v", err)
	}

	// Duplicate
	err = CreateCollection(userID, "My Books")
	if err != ErrCollectionExists {
		t.Fatalf("Expected ErrCollectionExists, got %v", err)
	}

	// Empty name
	err = CreateCollection(userID, "")
	if err != ErrEmptyCollectionName {
		t.Fatalf("Expected ErrEmptyCollectionName, got %v", err)
	}

	// Just whitespaces
	err = CreateCollection(userID, "   ")
	if err != ErrEmptyCollectionName {
		t.Fatalf("Expected ErrEmptyCollectionName for whitespaces, got %v", err)
	}

	// Trimming
	err = CreateCollection(userID, "  Trim Me  ")
	if err != nil {
		t.Fatalf("Expected nil for trimmed string, got error: %v", err)
	}

	// Invalid name (symbols)
	err = CreateCollection(userID, "My_Books!")
	if err != ErrInvalidCollectionName {
		t.Fatalf("Expected ErrInvalidCollectionName, got %v", err)
	}
}

func TestGetCollections(t *testing.T) {
	setupTestDB(t)
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")
	_ = CreateUser("user1@example.com", "User 1", "password")
	_ = CreateUser("user2@example.com", "User 2", "password")
	u1, _ := AuthenticateUser("user1@example.com", "password")
	u2, _ := AuthenticateUser("user2@example.com", "password")

	userID1 := u1.ID
	userID2 := u2.ID

	_ = CreateCollection(userID1, "Comics")
	_ = CreateCollection(userID1, "Movies")
	_ = CreateCollection(userID2, "Stamps")

	collections, err := GetCollections(userID1)
	if err != nil {
		t.Fatalf("Failed to get collections: %v", err)
	}

	if len(collections) != 2 {
		t.Fatalf("Expected 2 collections, got %d", len(collections))
	}

	// Verify order or contents loosely
	foundComics := false
	foundMovies := false
	for _, c := range collections {
		if c.Name == "Comics" {
			foundComics = true
		}
		if c.Name == "Movies" {
			foundMovies = true
		}
	}

	if !foundComics || !foundMovies {
		t.Errorf("Missing expected collections in results: %+v", collections)
	}

	// Test user 2
	collections2, _ := GetCollections(userID2)
	if len(collections2) != 1 {
		t.Errorf("Expected 1 collection for user2, got %d", len(collections2))
	}
}
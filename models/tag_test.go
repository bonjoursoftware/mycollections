package models

import (
	"mycollections/db"
	"testing"
)

func TestCreateTag(t *testing.T) {
	setupTestDB(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = CreateUser("taguser@example.com", "Tag User", "password")
	u, _ := AuthenticateUser("taguser@example.com", "password")
	_ = CreateCollection(u.ID, "Tag Collection")
	cols, _ := GetCollections(u.ID)
	c := cols[0]

	// Happy path
	err := CreateTag(c.ID, "Sci-Fi")
	if err != nil {
		t.Fatalf("Expected nil, got error: %v", err)
	}

	// Valid with hyphens and underscores
	err = CreateTag(c.ID, "Action_Movies-2023")
	if err != nil {
		t.Fatalf("Expected nil for valid tag, got error: %v", err)
	}

	// Duplicate
	err = CreateTag(c.ID, "Sci-Fi")
	if err != ErrTagExists {
		t.Fatalf("Expected ErrTagExists, got %v", err)
	}

	// Empty name
	err = CreateTag(c.ID, "")
	if err != ErrEmptyTagName {
		t.Fatalf("Expected ErrEmptyTagName, got %v", err)
	}

	// Just whitespaces
	err = CreateTag(c.ID, "   ")
	if err != ErrEmptyTagName {
		t.Fatalf("Expected ErrEmptyTagName for whitespaces, got %v", err)
	}

	// Invalid characters (spaces)
	err = CreateTag(c.ID, "Sci Fi")
	if err != ErrInvalidTagName {
		t.Fatalf("Expected ErrInvalidTagName for space, got %v", err)
	}

	// Trimming valid string
	err = CreateTag(c.ID, "  Trimmed-Tag  ")
	if err != nil {
		t.Fatalf("Expected nil for trimmed string, got error: %v", err)
	}

	// Invalid characters (symbols)
	err = CreateTag(c.ID, "SciFi!")
	if err != ErrInvalidTagName {
		t.Fatalf("Expected ErrInvalidTagName for symbols, got %v", err)
	}
}

func TestGetTags(t *testing.T) {
	setupTestDB(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = CreateUser("test@example.com", "User", "password")
	u, _ := AuthenticateUser("test@example.com", "password")
	_ = CreateCollection(u.ID, "Coll 1")
	cols, _ := GetCollections(u.ID)
	c := cols[0]

	_ = CreateTag(c.ID, "Tag1")
	_ = CreateTag(c.ID, "Tag2")

	tags, err := GetTags(c.ID)
	if err != nil {
		t.Fatalf("Failed to get tags: %v", err)
	}

	if len(tags) != 2 {
		t.Fatalf("Expected 2 tags, got %d", len(tags))
	}

	if tags[0].Name != "Tag1" && tags[1].Name != "Tag1" {
		t.Errorf("Missing Tag1")
	}
}

func TestUpdateTag(t *testing.T) {
	setupTestDB(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = CreateUser("test@example.com", "User", "password")
	u, _ := AuthenticateUser("test@example.com", "password")
	_ = CreateCollection(u.ID, "Coll")
	cols, _ := GetCollections(u.ID)
	c := cols[0]

	_ = CreateTag(c.ID, "OldName")
	tags, _ := GetTags(c.ID)
	tagID := tags[0].ID

	// Happy path
	err := UpdateTag(tagID, "NewName")
	if err != nil {
		t.Fatalf("Expected nil on update, got: %v", err)
	}

	// Duplicate
	_ = CreateTag(c.ID, "AnotherTag")
	err = UpdateTag(tagID, "AnotherTag")
	if err != ErrTagExists {
		t.Fatalf("Expected ErrTagExists, got %v", err)
	}

	// Invalid name
	err = UpdateTag(tagID, "Invalid Name")
	if err != ErrInvalidTagName {
		t.Fatalf("Expected ErrInvalidTagName, got %v", err)
	}
}

func TestDeleteTag(t *testing.T) {
	setupTestDB(t)
	_, _ = db.DB.Exec("DELETE FROM tags")
	_, _ = db.DB.Exec("DELETE FROM collections")
	_, _ = db.DB.Exec("DELETE FROM users")

	_ = CreateUser("test@example.com", "User", "password")
	u, _ := AuthenticateUser("test@example.com", "password")
	_ = CreateCollection(u.ID, "Coll")
	cols, _ := GetCollections(u.ID)
	c := cols[0]

	_ = CreateTag(c.ID, "ToDelete")
	tags, _ := GetTags(c.ID)
	tagID := tags[0].ID

	// Happy path
	err := DeleteTag(tagID)
	if err != nil {
		t.Fatalf("Expected nil on delete, got: %v", err)
	}

	tagsAfter, _ := GetTags(c.ID)
	if len(tagsAfter) != 0 {
		t.Fatalf("Expected 0 tags after deletion")
	}

	// Not found
	err = DeleteTag(9999)
	if err != ErrTagNotFound {
		t.Fatalf("Expected ErrTagNotFound, got %v", err)
	}
}

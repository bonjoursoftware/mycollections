package models

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"mycollections/db"
)

var (
	ErrTagExists     = errors.New("tag already exists in this collection")
	ErrEmptyTagName  = errors.New("tag name cannot be empty")
	ErrInvalidTagName = errors.New("tag name must be alphanumeric, hyphens, or underscores with no spaces")
	ErrTagNotFound   = errors.New("tag not found")
	tagNameRegex     = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
)

type Tag struct {
	ID           int
	CollectionID int
	Name         string
}

// validateTagName trims and checks if the tag name matches constraints.
func validateTagName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrEmptyTagName
	}
	if !tagNameRegex.MatchString(name) {
		return "", ErrInvalidTagName
	}
	return name, nil
}

// CreateTag inserts a new tag for a given collection, validating the name first.
func CreateTag(collectionID int, name string) error {
	name, err := validateTagName(name)
	if err != nil {
		return err
	}

	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM tags WHERE collection_id = ? AND name = ?", collectionID, name).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrTagExists
	}

	_, err = db.DB.Exec("INSERT INTO tags (collection_id, name) VALUES (?, ?)", collectionID, name)
	return err
}

// GetTags retrieves all tags for a specific collection.
func GetTags(collectionID int) ([]Tag, error) {
	rows, err := db.DB.Query("SELECT id, collection_id, name FROM tags WHERE collection_id = ?", collectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var t Tag
		if err := rows.Scan(&t.ID, &t.CollectionID, &t.Name); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

// UpdateTag modifies an existing tag name.
func UpdateTag(tagID int, newName string) error {
	newName, err := validateTagName(newName)
	if err != nil {
		return err
	}

	// Check if updating violates uniqueness for the same collection
	var collectionID int
	err = db.DB.QueryRow("SELECT collection_id FROM tags WHERE id = ?", tagID).Scan(&collectionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrTagNotFound
		}
		return err
	}

	var count int
	err = db.DB.QueryRow("SELECT COUNT(*) FROM tags WHERE collection_id = ? AND name = ? AND id != ?", collectionID, newName, tagID).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrTagExists
	}

	res, err := db.DB.Exec("UPDATE tags SET name = ? WHERE id = ?", newName, tagID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrTagNotFound
	}
	return nil
}

// DeleteTag removes a tag.
func DeleteTag(tagID int) error {
	res, err := db.DB.Exec("DELETE FROM tags WHERE id = ?", tagID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrTagNotFound
	}
	return nil
}

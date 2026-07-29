package models

import (
	"errors"
	"regexp"
	"strings"

	"mycollections/db"
)

var (
	ErrCollectionExists      = errors.New("collection already exists")
	ErrEmptyCollectionName   = errors.New("collection name cannot be empty")
	ErrInvalidCollectionName = errors.New("collection name must be alphanumeric and spaces")
	collectionNameRegex      = regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
)

type Collection struct {
	ID     int
	UserID int
	Name   string
}

// CreateCollection inserts a new collection for a given user.
func CreateCollection(userID int, name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyCollectionName
	}
	if !collectionNameRegex.MatchString(name) {
		return ErrInvalidCollectionName
	}

	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM collections WHERE user_id = ? AND name = ?", userID, name).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrCollectionExists
	}

	_, err = db.DB.Exec("INSERT INTO collections (user_id, name) VALUES (?, ?)", userID, name)
	return err
}

// GetCollections retrieves all collections belonging to a user.
func GetCollections(userID int) ([]Collection, error) {
	rows, err := db.DB.Query("SELECT id, user_id, name FROM collections WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []Collection
	for rows.Next() {
		var c Collection
		if err := rows.Scan(&c.ID, &c.UserID, &c.Name); err != nil {
			return nil, err
		}
		collections = append(collections, c)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return collections, nil
}

var ErrCollectionNotFound = errors.New("collection not found")

// GetCollectionByID retrieves a single collection belonging to a user.
func GetCollectionByID(userID, collectionID int) (*Collection, error) {
	var c Collection
	err := db.DB.QueryRow("SELECT id, user_id, name FROM collections WHERE id = ? AND user_id = ?", collectionID, userID).Scan(&c.ID, &c.UserID, &c.Name)
	if err != nil {
		if errors.Is(err, db.DB.QueryRow("").Err()) || err.Error() == "sql: no rows in result set" {
			return nil, ErrCollectionNotFound
		}
		return nil, err
	}
	return &c, nil
}
package models

import (
	"errors"
	"strings"

	"mycollections/db"
)

var (
	ErrEmptyItemName = errors.New("item name cannot be empty")
	ErrNoTagsForItem = errors.New("item must have at least one tag")
	ErrItemNotFound  = errors.New("item not found")
)

type Item struct {
	ID           int
	CollectionID int
	Name         string
	Description  string
	Tags         []Tag
}

// CreateItem inserts a new item and associates it with the given tags.
func CreateItem(collectionID int, name, description string, tagIDs []int) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyItemName
	}
	if len(tagIDs) == 0 {
		return ErrNoTagsForItem
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec("INSERT INTO items (collection_id, name, description) VALUES (?, ?, ?)", collectionID, name, description)
	if err != nil {
		return err
	}
	itemID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	for _, tagID := range tagIDs {
		_, err = tx.Exec("INSERT INTO item_tags (item_id, tag_id) VALUES (?, ?)", itemID, tagID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// UpdateItem modifies an existing item and its tag associations.
func UpdateItem(itemID int, name, description string, tagIDs []int) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return ErrEmptyItemName
	}
	if len(tagIDs) == 0 {
		return ErrNoTagsForItem
	}

	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.Exec("UPDATE items SET name = ?, description = ? WHERE id = ?", name, description, itemID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrItemNotFound
	}

	_, err = tx.Exec("DELETE FROM item_tags WHERE item_id = ?", itemID)
	if err != nil {
		return err
	}

	for _, tagID := range tagIDs {
		_, err = tx.Exec("INSERT INTO item_tags (item_id, tag_id) VALUES (?, ?)", itemID, tagID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// DeleteItem removes an item.
func DeleteItem(itemID int) error {
	res, err := db.DB.Exec("DELETE FROM items WHERE id = ?", itemID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return ErrItemNotFound
	}
	return nil
}

// GetItemByID retrieves a single item with its tags.
func GetItemByID(itemID int) (*Item, error) {
	var item Item
	err := db.DB.QueryRow("SELECT id, collection_id, name, description FROM items WHERE id = ?", itemID).Scan(&item.ID, &item.CollectionID, &item.Name, &item.Description)
	if err != nil {
		return nil, err
	}

	tagRows, err := db.DB.Query(`
		SELECT t.id, t.collection_id, t.name
		FROM tags t
		JOIN item_tags it ON t.id = it.tag_id
		WHERE it.item_id = ?
	`, itemID)
	if err != nil {
		return nil, err
	}
	defer tagRows.Close()

	var tags []Tag
	for tagRows.Next() {
		var t Tag
		if err := tagRows.Scan(&t.ID, &t.CollectionID, &t.Name); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	if err = tagRows.Err(); err != nil {
		return nil, err
	}
	item.Tags = tags

	return &item, nil
}

// GetItemsByTag retrieves all items associated with a specific tag.
func GetItemsByTag(tagID int) ([]Item, error) {
	rows, err := db.DB.Query(`
		SELECT i.id, i.collection_id, i.name, i.description 
		FROM items i
		JOIN item_tags it ON i.id = it.item_id
		WHERE it.tag_id = ?
	`, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.CollectionID, &item.Name, &item.Description); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Fetch tags for each item
	for i := range items {
		tagRows, err := db.DB.Query(`
			SELECT t.id, t.collection_id, t.name
			FROM tags t
			JOIN item_tags it ON t.id = it.tag_id
			WHERE it.item_id = ?
		`, items[i].ID)
		if err != nil {
			return nil, err
		}
		var tags []Tag
		for tagRows.Next() {
			var t Tag
			if err := tagRows.Scan(&t.ID, &t.CollectionID, &t.Name); err != nil {
				tagRows.Close()
				return nil, err
			}
			tags = append(tags, t)
		}
		tagRows.Close()
		if err = tagRows.Err(); err != nil {
			return nil, err
		}
		items[i].Tags = tags
	}

	return items, nil
}

package models

import (
	"testing"

	"mycollections/db"
)

func TestItemOperations(t *testing.T) {
	_ = db.InitDB(":memory:")
	defer db.CloseDB()

	// Setup user, collection, and tags
	_ = CreateUser("testitem@example.com", "Test", "pass")
	u, _ := AuthenticateUser("testitem@example.com", "pass")
	_ = CreateCollection(u.ID, "My Collection")
	cols, _ := GetCollections(u.ID)
	c := cols[0]
	
	_ = CreateTag(c.ID, "Tag1")
	_ = CreateTag(c.ID, "Tag2")
	tags, _ := GetTags(c.ID)
	t1, t2 := tags[0].ID, tags[1].ID

	// Create Item
	err := CreateItem(c.ID, "Item 1", "Desc 1", []int{t1})
	if err != nil {
		t.Fatalf("Failed to create item: %v", err)
	}

	// Empty name error
	err = CreateItem(c.ID, "  ", "Desc", []int{t1})
	if err != ErrEmptyItemName {
		t.Fatalf("Expected ErrEmptyItemName, got %v", err)
	}

	// No tags error
	err = CreateItem(c.ID, "Item 2", "Desc", []int{})
	if err != ErrNoTagsForItem {
		t.Fatalf("Expected ErrNoTagsForItem, got %v", err)
	}

	// Get Items by Tag
	items, err := GetItemsByTag(t1)
	if err != nil {
		t.Fatalf("Failed to get items: %v", err)
	}
	if len(items) != 1 || items[0].Name != "Item 1" {
		t.Fatalf("Expected 1 item 'Item 1', got %v", items)
	}
	if len(items[0].Tags) != 1 || items[0].Tags[0].ID != t1 {
		t.Fatalf("Expected item to have tag t1")
	}

	// Update Item
	err = UpdateItem(items[0].ID, "Updated Item", "Updated Desc", []int{t1, t2})
	if err != nil {
		t.Fatalf("Failed to update item: %v", err)
	}

	// Verify Update
	itemsT2, _ := GetItemsByTag(t2)
	if len(itemsT2) != 1 || itemsT2[0].Name != "Updated Item" {
		t.Fatalf("Expected update to reflect on t2")
	}
	if len(itemsT2[0].Tags) != 2 {
		t.Fatalf("Expected item to have 2 tags, got %d", len(itemsT2[0].Tags))
	}

	// Delete Item
	err = DeleteItem(items[0].ID)
	if err != nil {
		t.Fatalf("Failed to delete item: %v", err)
	}

	itemsT1, _ := GetItemsByTag(t1)
	if len(itemsT1) != 0 {
		t.Fatalf("Expected item to be deleted")
	}
}

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB initializes the database connection and creates required tables.
func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite", filepath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	err = createTables()
	if err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	log.Printf("Database initialized at %s\n", filepath)
	return nil
}

func createTables() error {
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		password_hash TEXT NOT NULL
	);`

	_, err := DB.Exec(usersTable)
	if err != nil {
		return err
	}

	collectionsTable := `
	CREATE TABLE IF NOT EXISTS collections (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		UNIQUE(user_id, name),
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(collectionsTable)
	if err != nil {
		return err
	}

	tagsTable := `
	CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		collection_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		UNIQUE(collection_id, name),
		FOREIGN KEY(collection_id) REFERENCES collections(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(tagsTable)
	if err != nil {
		return err
	}

	itemsTable := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		collection_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		FOREIGN KEY(collection_id) REFERENCES collections(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(itemsTable)
	if err != nil {
		return err
	}

	itemTagsTable := `
	CREATE TABLE IF NOT EXISTS item_tags (
		item_id INTEGER NOT NULL,
		tag_id INTEGER NOT NULL,
		PRIMARY KEY(item_id, tag_id),
		FOREIGN KEY(item_id) REFERENCES items(id) ON DELETE CASCADE,
		FOREIGN KEY(tag_id) REFERENCES tags(id) ON DELETE CASCADE
	);`

	_, err = DB.Exec(itemTagsTable)
	return err
}

// CloseDB closes the database connection gracefully.
func CloseDB() error {
	if DB != nil {
		log.Println("Closing database connection...")
		return DB.Close()
	}
	return nil
}

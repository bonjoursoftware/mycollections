package models

import (
	"database/sql"
	"errors"

	"mycollections/db"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailExists = errors.New("email already exists")
)

type User struct {
	ID           int
	Email        string
	Name         string
	PasswordHash string
}

// CreateUser hashes the password and inserts a new user into the database.
func CreateUser(email, name, password string) error {
	// Check if email exists
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("INSERT INTO users (email, name, password_hash) VALUES (?, ?, ?)", email, name, string(hashedPassword))
	return err
}

// AuthenticateUser verifies the email and password, returning the User object if successful.
func AuthenticateUser(email, password string) (*User, error) {
	u := &User{}
	err := db.DB.QueryRow("SELECT id, email, name, password_hash FROM users WHERE email = ?", email).Scan(&u.ID, &u.Email, &u.Name, &u.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return u, nil
}

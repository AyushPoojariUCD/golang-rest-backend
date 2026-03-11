package models

import (
	"go-rest-backend/db"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Save user (Signup)
func (u *User) Save() error {

	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		u.Email,
		u.Password,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = int(id)

	return nil
}

// Get user by email (Login)
func GetUserByEmail(email string) (*User, error) {

	query := `
	SELECT id, email, password
	FROM users
	WHERE email = ?
	`

	row := db.DB.QueryRow(query, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Get user by ID
func GetUserByID(id string) (*User, error) {

	query := `
	SELECT id, email, password
	FROM users
	WHERE id = ?
	`

	row := db.DB.QueryRow(query, id)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
package models

import (
	"errors"
	"go-rest-backend/db"
	"go-rest-backend/utils"
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

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(
		u.Email,
		hashedPassword,
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

// Get users
func GetAllUsers() ([]User, error) {

	query := `
	SELECT id, email, password
	FROM users
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
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

// Update user
func (u *User) Update() error {

	query := `
	UPDATE users
	SET email = ?, password = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		u.Email,
		u.Password,
		u.ID,
	)

	return err
}

// Delete user
func DeleteUser(id string) error {

	query := `DELETE FROM users WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (u *User) ValidateCredential() error {

	query := `
	SELECT id, password
	FROM users
	WHERE email = ?
	`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
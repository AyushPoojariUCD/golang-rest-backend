package utils

import "golang.org/x/crypto/bcrypt"

// Hash password before saving to DB
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Compare password with stored hash
func CheckPasswordHash(password string, hash string) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err == nil
}
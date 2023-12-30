package models

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

type NewUser struct {
	Email    string
	Password string
}

func (us *UserService) Create(email, password string) (*User, error) {

	if !isValidEmail(email) {
		return nil, fmt.Errorf("invalid email format: %s", email)
	}

	email = strings.ToLower(email)

	hashedBytes, err := generateHashBytesForPassword(password)
	if err != nil {
		return nil, fmt.Errorf("Create user: %w", err)
	}
	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(`INSERT  INTO users (email, password_hash) VALUES ($1, $2) RETURNING id`, email, passwordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil
}

func (us UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)

	if !isValidEmail(email) {
		return nil, fmt.Errorf("invalid email format: %s", email)
	}

	user := User{
		Email: email,
	}

	row := us.DB.QueryRow(`SELECT id, password_hash FROM users WHERE email=$1`, email)
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	prehashedPassword := sha256.Sum256([]byte(password))
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), prehashedPassword[:])
	if err != nil {
		return nil, fmt.Errorf("authetnicated: %w", err)
	}
	return &user, nil
}

func isValidEmail(email string) bool {
	// check for enmpty string
	if email == "" {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}

	localPart, domainPart := parts[0], parts[1]

	if len(localPart) == 0 || len(domainPart) == 0 {
		return false
	}

	if len(localPart) > 64 {
		return false
	}

	if len(domainPart) > 255 {
		return false
	}

	// Check for invalid characters in local part
	if hasInvalidChars(localPart, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-_") {
		return false
	}
	// check for invalid characters in domain part
	if hasInvalidChars(domainPart, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.-_") {
		return false
	}

	return true
}

func hasInvalidChars(s, validChars string) bool {
	// check if any character in the string is NOT a valid character
	for _, char := range s {
		if !strings.ContainsRune(validChars, char) {
			return true
		}
	}
	return false
}

func generateHashBytesForPassword(passwordString string) ([]byte, error) {
	// Prehash with SHA-256
	prehashedPassword := sha256.Sum256([]byte(passwordString))

	hashedBytes, err := bcrypt.GenerateFromPassword(prehashedPassword[:], bcrypt.DefaultCost)

	return hashedBytes, err
}

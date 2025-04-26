package model

import (
	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model             
    Username string
    PasswordHash string
    Role     string
    FullName string
    Email    string
}

func NewUser(email, password, fullName, username, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		PasswordHash: string(hashedPassword),
		FullName: fullName,
        Username: username,
        Role:     role,
	}, nil
}



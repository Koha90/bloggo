// Package types - contain the types for application.
package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User - object of user/account of application bloggo.
type User struct {
	ID                uint      `json:"id"`
	Username          string    `json:"username"`
	FirstName         string    `json:"first_name,omitempty"`
	LastName          string    `json:"last_name,omitempty"`
	Role              string    `json:"role"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

// CreateUserRequest - fields for new user.
type CreateUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

// ValidatePassword - say for himself.
func (u *User) ValidatePassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw)) == nil
}

// NewUser - with password.
func NewUser(username, firstName, lastName, password string) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Username:          username,
		FirstName:         firstName,
		LastName:          lastName,
		Role:              "user",
		EncryptedPassword: string(encpw),
		CreatedAt:         time.Now().UTC(),
		UpdatedAt:         time.Now().UTC(),
	}, nil
}

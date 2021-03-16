package users

import (
	"strings"

	"github.com/agusluques/bookstore_utils-go/rest_errors"
)

const (
	// StatusActive enum
	StatusActive = "active"
)

// User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users is a slice of User
type Users []User

// Validate an user
func (u *User) Validate() *rest_errors.RestError {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	u.Password = strings.TrimSpace(u.Password)
	if u.Password == "" {
		return rest_errors.NewBadRequestError("invalid password")
	}

	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return rest_errors.NewBadRequestError("invalid email address")
	}

	return nil
}

package users

import (
	"strings"

	"github.com/agusluques/bookstore_users-api/utils/errors"
)

type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (u *User) Validate() *errors.RestError {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))

	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}

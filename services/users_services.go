package services

import (
	"github.com/agusluques/bookstore_users-api/domain/users"
	utils "github.com/agusluques/bookstore_users-api/utils/date_utils"
	"github.com/agusluques/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := users.User{
		ID: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = utils.GetNowDBString()
	user.Status = users.StatusActive

	if err := user.Save(); err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(isPartial bool, user *users.User) (*users.User, *errors.RestError) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			currentUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			currentUser.LastName = user.LastName
		}
		if user.Email != "" {
			currentUser.Email = user.Email
		}
	} else {
		currentUser.FirstName = user.FirstName
		currentUser.LastName = user.LastName
		currentUser.Email = user.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}

	return currentUser, nil
}

func DeleteUser(userId int64) *errors.RestError {
	user, err := GetUser(userId)
	if err != nil {
		return err
	}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func Search(status string) (*[]users.User, *errors.RestError) {
	users, err := users.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

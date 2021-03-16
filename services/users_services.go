package services

import (
	"github.com/agusluques/bookstore_users-api/domain/users"
	"github.com/agusluques/bookstore_users-api/utils/crypto_utils"
	utils "github.com/agusluques/bookstore_users-api/utils/date_utils"
	"github.com/agusluques/bookstore_utils-go/rest_errors"
)

// UsersService variable
var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

type usersServiceInterface interface {
	Get(int64) (*users.User, *rest_errors.RestError)
	Create(*users.User) (*users.User, *rest_errors.RestError)
	Update(bool, *users.User) (*users.User, *rest_errors.RestError)
	DeleteUser(int64) *rest_errors.RestError
	Search(string) (users.Users, *rest_errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestError)
}

func (s *usersService) Get(userID int64) (*users.User, *rest_errors.RestError) {
	user := users.User{
		ID: userID,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) Create(user *users.User) (*users.User, *rest_errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = utils.GetNowDBString()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *usersService) Update(isPartial bool, user *users.User) (*users.User, *rest_errors.RestError) {
	currentUser, err := UsersService.Get(user.ID)
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

func (s *usersService) DeleteUser(userID int64) *rest_errors.RestError {
	user, err := UsersService.Get(userID)
	if err != nil {
		return err
	}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

func (s *usersService) Search(status string) (users.Users, *rest_errors.RestError) {
	users, err := users.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *usersService) LoginUser(req users.LoginRequest) (*users.User, *rest_errors.RestError) {
	dao := &users.User{
		Email:    req.Email,
		Password: crypto_utils.GetMd5(req.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}

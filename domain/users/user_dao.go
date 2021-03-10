package users

import (
	"strings"

	"github.com/agusluques/bookstore_users-api/datasources/mysql/users_db"
	"github.com/agusluques/bookstore_users-api/logger"
	"github.com/agusluques/bookstore_users-api/utils/errors"
	"github.com/agusluques/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?"
)

// Get an user
func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error while trying to get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

// Save an user
func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error while trying to prepare user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error while trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error while trying to get user last id", err)
		return errors.NewInternalServerError("database error")
	}

	user.ID = userID

	return user.Get()
}

// Update an user
func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error while trying to prepare user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if updateErr != nil {
		logger.Error("error while trying to update user", updateErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

// Delete an user
func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error while trying to prepare user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.ID)
	if deleteErr != nil {
		logger.Error("error while trying to delete user", deleteErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

// FindByStatus an user
func FindByStatus(status string) ([]User, *errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error while trying to prepare user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error while trying to find users", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	var results = make([]User, 0)
	for rows.Next() {
		var user User

		if getErr := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
			logger.Error("error while trying to get user", getErr)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	return results, nil
}

// FindByEmailAndPassword return an user
func (user *User) FindByEmailAndPassword() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error while trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		if strings.Contains(getErr.Error(), mysql_utils.ErrorNoRows) {
			return errors.NewNotFoundError("invalid users credentials")
		}
		logger.Error("error while trying to get user by email and password", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

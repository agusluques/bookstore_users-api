package mysql_utils

import (
	"strings"

	"github.com/agusluques/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = 1062
	ErrorNoRows      = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case indexUniqueEmail:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}

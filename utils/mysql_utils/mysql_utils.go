package mysql_utils

import (
	"strings"

	"github.com/agusluques/bookstore_utils-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	indexUniqueEmail = 1062
	ErrorNoRows      = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database response", sqlErr)
	}

	switch sqlErr.Number {
	case indexUniqueEmail:
		return rest_errors.NewBadRequestError("invalid data")
	}

	return rest_errors.NewInternalServerError("error processing request", sqlErr)
}

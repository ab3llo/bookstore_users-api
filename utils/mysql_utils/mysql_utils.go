package mysql_utils

import (
	"strings"

	"github.com/ab3llo/bookstore_users-api/logger"
	"github.com/ab3llo/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNotFound = "record not found"
)

//ParseError return restError
func ParseError(err error) *errors.RestError {
	sqlError, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNotFound) {
			logger.Error("no record found matching the given id", err)
			return errors.NewNotFoundError("no record matching the given id")
		}
		logger.Error("database error", err)
		return errors.NewInternalServerError("database error")
	}

	switch sqlError.Number {
	case 1062:
		logger.Error("Email address already exists", err)
		return errors.NewBadRequestError("invalid data")
	}
	logger.Error("Error when processing request", err)
	return errors.NewInternalServerError("error when processing request")
}

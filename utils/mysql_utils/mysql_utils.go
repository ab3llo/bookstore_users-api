package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/ab3llo/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNotFound = "record not found"
)

//ParseError return restError
func ParseError(err error) *errors.RestError {
	sqlError, ok := err.(*mysql.MySQLError)
	fmt.Print("OK: ", ok)
	if !ok {
		fmt.Println("This ERROR: " + err.Error())
		if strings.Contains(err.Error(), errNotFound) {
			return errors.NewNotFoundError("no record matching the given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlError.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error when processing request")
}

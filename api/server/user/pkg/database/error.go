package database

import (
	"errors"
	"fmt"

	ce "github.com/calmato/gran-book/api/server/user/pkg/errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func ToDBError(err error) error {
	if err == nil {
		return nil
	}

	//nolint:gocritic
	switch err.(type) {
	case *mysql.MySQLError:
		return newDBError(err.Error(), ce.ErrInternal)
	}

	switch {
	case errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrInvalidField),
		errors.Is(err, gorm.ErrInvalidTransaction),
		errors.Is(err, gorm.ErrInvalidValue),
		errors.Is(err, gorm.ErrInvalidValueOfLength),
		errors.Is(err, gorm.ErrMissingWhereClause),
		errors.Is(err, gorm.ErrModelValueRequired),
		errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return newDBError(err.Error(), ce.ErrBadRequest)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return newDBError(err.Error(), ce.ErrNotFound)
	case errors.Is(err, gorm.ErrNotImplemented):
		return newDBError(err.Error(), ce.ErrServiceUnabailable)
	case errors.Is(err, gorm.ErrDryRunModeUnsupported),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrUnsupportedRelation):
		return newDBError(err.Error(), ce.ErrInternal)
	default:
		return newDBError(err.Error(), ce.ErrUnknown)
	}
}

func newDBError(str string, err error) error {
	return fmt.Errorf("database: %s: %w", str, err)
}

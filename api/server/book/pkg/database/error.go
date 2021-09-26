package database

import (
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("database: invalid transaction")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("database: not implemented")
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = errors.New("database: record not found")
	// ErrInternal internal error
	ErrInternal = errors.New("database: internal error")
	// ErrUnknown unknown error
	ErrUnknown = errors.New("database: unknown error")
)

func ToDBError(err error) error {
	if err == nil {
		return nil
	}

	switch err.(type) {
	case *mysql.MySQLError:
		return fmt.Errorf("%v: %w", err, ErrInternal)
	}

	switch {
	case
		errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrInvalidField),
		errors.Is(err, gorm.ErrInvalidTransaction),
		errors.Is(err, gorm.ErrInvalidValue),
		errors.Is(err, gorm.ErrInvalidValueOfLength),
		errors.Is(err, gorm.ErrMissingWhereClause),
		errors.Is(err, gorm.ErrModelValueRequired),
		errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return fmt.Errorf("%v: %w", ErrInvalidTransaction, err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return fmt.Errorf("%v: %w", ErrRecordNotFound, err)
	case
		errors.Is(err, gorm.ErrNotImplemented),
		errors.Is(err, gorm.ErrUnsupportedDriver):
		return fmt.Errorf("%v: %w", ErrNotImplemented, err)
	case
		errors.Is(err, gorm.ErrDryRunModeUnsupported),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrUnsupportedRelation):
		return fmt.Errorf("%v: %w", ErrInternal, err)
	default:
		return fmt.Errorf("%v: %w", ErrUnknown, err)
	}
}

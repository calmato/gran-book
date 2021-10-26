package exception

import (
	"errors"

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
		return ErrInDatastore.New(err)
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
		return ErrInvalidArgument.New(err)
	case errors.Is(err, gorm.ErrRecordNotFound):
		return ErrNotFound.New(err)
	case errors.Is(err, gorm.ErrNotImplemented):
		return ErrNotImplemented.New(err)
	case errors.Is(err, gorm.ErrDryRunModeUnsupported),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrUnsupportedRelation):
		return ErrInternal.New(err)
	default:
		return ErrUnknown.New(err)
	}
}

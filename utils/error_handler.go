package utils

import (
	"errors"
	"go-modules-api/internal/exceptions"
	"gorm.io/gorm"
)

// HandleDBError handles database errors and returns an APIException
func HandleDBError(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return exceptions.NotFound("Record not found", nil)

	case errors.Is(err, gorm.ErrDuplicatedKey):
		return exceptions.Conflict("A record with this value already exists", nil)

	case errors.Is(err, gorm.ErrForeignKeyViolated):
		return exceptions.BadRequest("Foreign key constraint violation", nil)

	case errors.Is(err, gorm.ErrInvalidTransaction):
		return exceptions.BadRequest("Invalid transaction", nil)

	case errors.Is(err, gorm.ErrCheckConstraintViolated):
		return exceptions.BadRequest("Check constraint violation", nil)

	default:
		return exceptions.InternalServerError("A database error occurred", nil)
	}
}

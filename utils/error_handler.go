package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go-modules-api/internal/exceptions"
	"gorm.io/gorm"
)

// HandleDBError handles database errors and returns an APIException
func HandleDBError(err error, field string, value interface{}, message ...string) error {
	if err == nil {
		return nil
	}

	defaultMessages := map[error]string{
		gorm.ErrRecordNotFound:          "Record not found",
		gorm.ErrDuplicatedKey:           "A record with this value already exists",
		gorm.ErrForeignKeyViolated:      "Foreign key constraint violation",
		gorm.ErrInvalidTransaction:      "Invalid transaction",
		gorm.ErrCheckConstraintViolated: "Check constraint violation",
	}

	customMessage := ""
	if len(message) > 0 && message[0] != "" {
		customMessage = message[0]
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg := customMessage
		if msg == "" {
			msg = defaultMessages[gorm.ErrRecordNotFound]
		}
		return exceptions.NotFound(msg, fiber.Map{"field": field, "value": value})
	}

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		msg := customMessage
		if msg == "" {
			msg = defaultMessages[gorm.ErrDuplicatedKey]
		}
		return exceptions.Conflict(msg, fiber.Map{"field": field, "value": value})
	}

	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		msg := customMessage
		if msg == "" {
			msg = defaultMessages[gorm.ErrForeignKeyViolated]
		}
		return exceptions.BadRequest(msg, nil)
	}

	if errors.Is(err, gorm.ErrInvalidTransaction) {
		msg := customMessage
		if msg == "" {
			msg = defaultMessages[gorm.ErrInvalidTransaction]
		}
		return exceptions.BadRequest(msg, nil)
	}

	if errors.Is(err, gorm.ErrCheckConstraintViolated) {
		msg := customMessage
		if msg == "" {
			msg = defaultMessages[gorm.ErrCheckConstraintViolated]
		}
		return exceptions.BadRequest(msg, nil)
	}

	msg := "A database error occurred"
	if customMessage != "" {
		msg = customMessage
	}
	return exceptions.InternalServerError(msg, err.Error())
}

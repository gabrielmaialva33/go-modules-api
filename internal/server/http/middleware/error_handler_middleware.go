package middleware

import (
	"errors"

	"go-modules-api/internal/exceptions"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	var apiErr *exceptions.APIException
	if errors.As(err, &apiErr) {
		return apiErr.Response(c)
	}

	return exceptions.InternalServerError("An unexpected error occurred", err.Error()).Response(c)
}

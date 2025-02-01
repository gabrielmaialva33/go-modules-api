package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if strings.Contains(err.Error(), "invalid") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   "bad_request",
			"message": "Invalid request data.",
		})
	}

	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  fiber.StatusConflict,
			"error":   "duplicate_key",
			"message": "A unique field value already exists.",
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  fiber.StatusInternalServerError,
		"error":   "internal_server_error",
		"message": "An unexpected error occurred. Please try again later.",
	})
}

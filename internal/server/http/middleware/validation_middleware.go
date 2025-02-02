package middleware

import "github.com/gofiber/fiber/v2"

func ValidationMiddleware(c *fiber.Ctx) error {
	body := new(map[string]interface{})

	// Parse Body
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON format",
		})
	}

	return c.Next()
}

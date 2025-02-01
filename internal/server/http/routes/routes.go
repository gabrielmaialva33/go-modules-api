package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-modules-api/internal/server/container"
	"go.uber.org/zap"
)

func SetupRoutes(app *fiber.App, log *zap.Logger, container *container.AppContainer) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/redoc.html")
	})

	HubClientsRoutes(app, container.Handlers.HubClientHandler)

}

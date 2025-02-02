package routes

import (
	"go-modules-api/internal/server/http/handlers"

	"github.com/gofiber/fiber/v2"
)

// HubClientsRoutes defines routes for hub clients
func HubClientsRoutes(app *fiber.App, hubClientHandler *handlers.HubClientHandler) {
	api := app.Group("/api")

	// Hub Clients Routes
	hubClients := api.Group("/hub_clients")

	hubClients.Get("/paginate", hubClientHandler.PaginateHubClients)
	hubClients.Get("/", hubClientHandler.GetAllHubClients)
	hubClients.Post("/", hubClientHandler.CreateHubClient)
	hubClients.Put("/:id", hubClientHandler.UpdateHubClient)
	hubClients.Delete("/:id", hubClientHandler.DeleteHubClient)

	hubClients.Get("/:id", hubClientHandler.GetHubClientByID)
}

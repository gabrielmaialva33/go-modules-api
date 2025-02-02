package routes

import (
	"go-modules-api/internal/server/http/handlers"

	"github.com/gofiber/fiber/v2"
)

// RoleRoutes defines routes for roles
func RoleRoutes(app *fiber.App, roleHandler *handlers.RoleHandler) {
	api := app.Group("/api")

	// Roles Routes
	roles := api.Group("/roles")

	roles.Get("/paginate", roleHandler.PaginateRoles)
	roles.Get("/", roleHandler.ListRoles)
	roles.Post("/", roleHandler.CreateRole)
	roles.Put("/:id", roleHandler.UpdateRole)
	roles.Delete("/:id", roleHandler.SoftDeleteRole)

	roles.Get("/:id", roleHandler.GetRoleByID)
}

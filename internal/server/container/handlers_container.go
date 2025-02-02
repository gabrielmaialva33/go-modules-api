package container

import (
	"go-modules-api/internal/server/http/handlers"
)

type HandlersContainer struct {
	HubClientHandler *handlers.HubClientHandler
	RoleHandler      *handlers.RoleHandler
}

func NewHandlersContainer(services *ServicesContainer) *HandlersContainer {
	hubClientHandler := handlers.NewHubClientHandler(services.HubClientService)
	roleHandler := handlers.NewRoleHandler(services.RoleService)

	return &HandlersContainer{
		HubClientHandler: hubClientHandler,
		RoleHandler:      roleHandler,
	}
}

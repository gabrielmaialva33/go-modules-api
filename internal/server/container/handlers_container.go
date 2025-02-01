package container

import (
	"go-modules-api/internal/server/http/handlers"
)

type HandlersContainer struct {
	HubClientHandler *handlers.HubClientHandler
}

func NewHandlersContainer(services *ServicesContainer) *HandlersContainer {
	hubClientHandler := handlers.NewHubClientHandler(services.HubClientService)

	return &HandlersContainer{
		HubClientHandler: hubClientHandler,
	}
}

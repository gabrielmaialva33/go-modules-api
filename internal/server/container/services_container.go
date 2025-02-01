package container

import (
	"go-modules-api/internal/services"
)

type ServicesContainer struct {
	HubClientService services.HubClientService
}

func NewServicesContainer(repositories *RepositoriesContainer) *ServicesContainer {
	hubClientService := services.NewHubClientService(repositories.HubClientRepository)

	return &ServicesContainer{
		HubClientService: hubClientService,
	}
}

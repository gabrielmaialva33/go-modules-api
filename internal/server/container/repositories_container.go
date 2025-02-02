package container

import (
	"go-modules-api/config"
	"go-modules-api/internal/repositories"
)

type RepositoriesContainer struct {
	HubClientRepository repositories.HubClientRepository
}

func NewRepositoriesContainer() *RepositoriesContainer {
	hubClientRepository := repositories.NewHubClientRepository(config.DB)

	return &RepositoriesContainer{
		HubClientRepository: hubClientRepository,
	}
}

package container

import (
	"go-modules-api/internal/repositories"
)

type RepositoriesContainer struct {
	HubClientRepository repositories.HubClientRepository
}

func NewRepositoriesContainer() *RepositoriesContainer {
	hubClientRepository := repositories.NewHubClientRepository()

	return &RepositoriesContainer{
		HubClientRepository: hubClientRepository,
	}
}

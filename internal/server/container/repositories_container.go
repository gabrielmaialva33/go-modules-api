package container

import (
	"go-modules-api/config"
	"go-modules-api/internal/repositories"
)

type RepositoriesContainer struct {
	HubClientRepository repositories.HubClientRepository
	RoleRepository      repositories.RoleRepository
}

func NewRepositoriesContainer() *RepositoriesContainer {
	hubClientRepository := repositories.NewHubClientRepository(config.DB)
	roleRepository := repositories.NewRoleRepository(config.DB)

	return &RepositoriesContainer{
		HubClientRepository: hubClientRepository,
		RoleRepository:      roleRepository,
	}
}

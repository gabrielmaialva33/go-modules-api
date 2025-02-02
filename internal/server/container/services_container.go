package container

import (
	"go-modules-api/internal/services"
)

type ServicesContainer struct {
	HubClientService services.HubClientService
	RoleService      services.RoleService
}

func NewServicesContainer(repositories *RepositoriesContainer) *ServicesContainer {
	hubClientService := services.NewHubClientService(repositories.HubClientRepository)
	roleService := services.NewRoleService(repositories.RoleRepository)

	return &ServicesContainer{
		HubClientService: hubClientService,
		RoleService:      roleService,
	}
}

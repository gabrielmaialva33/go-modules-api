package container

type AppContainer struct {
	Repositories *RepositoriesContainer
	Services     *ServicesContainer
	Handlers     *HandlersContainer
}

func NewAppContainer() *AppContainer {
	repositories := NewRepositoriesContainer()
	services := NewServicesContainer(repositories)
	handlers := NewHandlersContainer(services)

	return &AppContainer{
		Repositories: repositories,
		Services:     services,
		Handlers:     handlers,
	}
}

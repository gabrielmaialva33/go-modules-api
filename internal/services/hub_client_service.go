package services

import (
	"go-modules-api/internal/dto"
	"go-modules-api/internal/models"
	"go-modules-api/internal/repositories"
	"go-modules-api/utils"
)

// HubClientService defines business logic for hub clients
type HubClientService interface {
	PaginateHubClients(params dto.PaginatedHubClientDTO) ([]models.HubClient, int64, error)
	ListHubClients(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error)
	GetHubClientByID(id uint) (*models.HubClient, error)
	CreateHubClient(hubClient *models.HubClient) error
	UpdateHubClient(hubClient *models.HubClient) error
	DeleteHubClient(id uint) error
}

type hubClientService struct {
	repo repositories.HubClientRepository
}

func NewHubClientService(repo repositories.HubClientRepository) HubClientService {
	return &hubClientService{repo: repo}
}

// PaginateHubClients retrieves paginated hub clients
func (s *hubClientService) PaginateHubClients(params dto.PaginatedHubClientDTO) ([]models.HubClient, int64, error) {
	clients, total, err := s.repo.Pagination(
		params.Search,
		params.Active,
		params.SortField,
		params.SortOrder,
		params.Page,
		params.PageSize,
	)
	return clients, total, utils.HandleDBError(err)
}

// ListHubClients returns all hub clients with filtering and sorting
func (s *hubClientService) ListHubClients(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error) {
	clients, err := s.repo.GetAll(search, active, sortField, sortOrder)
	return clients, utils.HandleDBError(err)
}

// GetHubClientByID retrieves a hub client by ID
func (s *hubClientService) GetHubClientByID(id uint) (*models.HubClient, error) {
	client, err := s.repo.GetByID(id)
	return client, utils.HandleDBError(err)
}

// CreateHubClient creates a new hub client
func (s *hubClientService) CreateHubClient(hubClient *models.HubClient) error {
	return utils.HandleDBError(s.repo.Create(hubClient))
}

// UpdateHubClient updates an existing hub client
func (s *hubClientService) UpdateHubClient(hubClient *models.HubClient) error {
	return utils.HandleDBError(s.repo.Update(hubClient))
}

// DeleteHubClient removes a hub client
func (s *hubClientService) DeleteHubClient(id uint) error {
	return utils.HandleDBError(s.repo.Delete(id))
}

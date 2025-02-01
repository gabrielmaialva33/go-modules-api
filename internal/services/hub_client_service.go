package services

import (
	"go-modules-api/internal/models"
	"go-modules-api/internal/repositories"
)

// HubClientService defines business logic for hub clients
type HubClientService interface {
	GetAllHubClients() ([]models.HubClient, error)
	GetHubClientByID(id uint) (*models.HubClient, error)
	CreateHubClient(hubClient *models.HubClient) error
	UpdateHubClient(hubClient *models.HubClient) error
	DeleteHubClient(id uint) error
}

// hubClientService implements HubClientService
type hubClientService struct {
	repo repositories.HubClientRepository
}

// NewHubClientService creates a new service instance
func NewHubClientService(repo repositories.HubClientRepository) HubClientService {
	return &hubClientService{repo: repo}
}

// GetAllHubClients returns all hub clients
func (s *hubClientService) GetAllHubClients() ([]models.HubClient, error) {
	return s.repo.GetAll()
}

// GetHubClientByID retrieves a hub client by ID
func (s *hubClientService) GetHubClientByID(id uint) (*models.HubClient, error) {
	return s.repo.GetByID(id)
}

// CreateHubClient creates a new hub client
func (s *hubClientService) CreateHubClient(hubClient *models.HubClient) error {
	return s.repo.Create(hubClient)
}

// UpdateHubClient updates an existing hub client
func (s *hubClientService) UpdateHubClient(hubClient *models.HubClient) error {
	return s.repo.Update(hubClient)
}

// DeleteHubClient removes a hub client
func (s *hubClientService) DeleteHubClient(id uint) error {
	return s.repo.Delete(id)
}

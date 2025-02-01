package services

import (
	"go-modules-api/internal/models"
	"go-modules-api/internal/repositories"
	"go-modules-api/utils"
)

// HubClientService defines business logic for hub clients
type HubClientService interface {
	GetAllHubClients() ([]models.HubClient, error)
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

// GetAllHubClients returns all hub clients
func (s *hubClientService) GetAllHubClients() ([]models.HubClient, error) {
	clients, err := s.repo.GetAll()
	return clients, utils.HandleDBError(err, "", nil)
}

// GetHubClientByID retrieves a hub client by ID
func (s *hubClientService) GetHubClientByID(id uint) (*models.HubClient, error) {
	client, err := s.repo.GetByID(id)
	return client, utils.HandleDBError(err, "id", id, "Hub client not found")
}

// CreateHubClient creates a new hub client
func (s *hubClientService) CreateHubClient(hubClient *models.HubClient) error {
	return utils.HandleDBError(s.repo.Create(hubClient), "external_id", hubClient.ExternalID)
}

// UpdateHubClient updates an existing hub client
func (s *hubClientService) UpdateHubClient(hubClient *models.HubClient) error {
	return utils.HandleDBError(s.repo.Update(hubClient), "id", hubClient.ID, "Could not update hub client")
}

// DeleteHubClient removes a hub client
func (s *hubClientService) DeleteHubClient(id uint) error {
	return utils.HandleDBError(s.repo.Delete(id), "id", id, "Could not delete hub client")
}

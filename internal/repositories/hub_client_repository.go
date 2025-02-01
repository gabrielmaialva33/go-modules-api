package repositories

import (
	"go-modules-api/config"
	"go-modules-api/internal/models"
)

// HubClientRepository defines the interface for database operations
type HubClientRepository interface {
	GetAll() ([]models.HubClient, error)
	GetByID(id uint) (*models.HubClient, error)
	Create(hubClient *models.HubClient) error
	Update(hubClient *models.HubClient) error
	Delete(id uint) error
}

// hubClientRepository implements HubClientRepository interface
type hubClientRepository struct{}

// NewHubClientRepository creates a new instance of HubClientRepository
func NewHubClientRepository() HubClientRepository {
	return &hubClientRepository{}
}

// GetAll retrieves all hub clients from the database
func (r *hubClientRepository) GetAll() ([]models.HubClient, error) {
	var clients []models.HubClient
	err := config.DB.Find(&clients).Error
	return clients, err
}

// GetByID retrieves a single hub client by ID
func (r *hubClientRepository) GetByID(id uint) (*models.HubClient, error) {
	var client models.HubClient
	err := config.DB.First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// Create inserts a new hub client into the database
func (r *hubClientRepository) Create(hubClient *models.HubClient) error {
	return config.DB.Create(hubClient).Error
}

// Update modifies an existing hub client
func (r *hubClientRepository) Update(hubClient *models.HubClient) error {
	return config.DB.Save(hubClient).Error
}

// Delete removes a hub client from the database
func (r *hubClientRepository) Delete(id uint) error {
	return config.DB.Delete(&models.HubClient{}, id).Error
}

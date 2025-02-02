package repositories

import (
	"go-modules-api/config"
	"go-modules-api/internal/models"
)

// HubClientRepository defines the interface for database operations
type HubClientRepository interface {
	Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.HubClient, int64, error)
	GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error)
	GetByID(id uint) (*models.HubClient, error)
	Create(hubClient *models.HubClient) error
	Update(hubClient *models.HubClient) error
	Delete(id uint) error
}

type hubClientRepository struct{}

// NewHubClientRepository creates a new instance of HubClientRepository
func NewHubClientRepository() HubClientRepository {
	return &hubClientRepository{}
}

// Pagination retrieves paginated hub clients from the database
func (r *hubClientRepository) Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.HubClient, int64, error) {
	var clients []models.HubClient
	var total int64

	db := config.DB.Model(&models.HubClient{})

	if search != "" {
		db = db.Where("name ILIKE ?", "%"+search+"%")
	}

	if active != nil {
		db = db.Where("active = ?", *active)
	}

	db.Count(&total)

	if sortField != "" {
		if sortOrder != "desc" {
			sortOrder = "asc"
		}
		db = db.Order(sortField + " " + sortOrder)
	}

	offset := (page - 1) * pageSize
	err := db.Offset(offset).Limit(pageSize).Find(&clients).Error

	return clients, total, err
}

// GetAll retrieves all hub clients from the database with filtering and sorting
func (r *hubClientRepository) GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error) {
	var clients []models.HubClient
	db := config.DB

	if search != "" {
		db = db.Where("name ILIKE ?", "%"+search+"%")
	}

	if active != nil {
		db = db.Where("active = ?", *active)
	}

	if sortField != "" {
		if sortOrder != "desc" {
			sortOrder = "asc"
		}
		db = db.Order(sortField + " " + sortOrder)
	}

	err := db.Find(&clients).Error
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
	return config.DB.Where("id = ?", hubClient.ID).Updates(hubClient).Error
}

// Delete removes a hub client from the database
func (r *hubClientRepository) Delete(id uint) error {
	return config.DB.Delete(&models.HubClient{}, id).Error
}

package repositories

import (
	"go-modules-api/internal/models"
	"gorm.io/gorm"
)

// HubClientRepository defines the interface for database operations specific to HubClient.
type HubClientRepository interface {
	BaseRepositoryInterface[*models.HubClient]
	Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.HubClient, int64, error)
	GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error)
}

type hubClientRepository struct {
	base *BaseRepository[*models.HubClient]
	db   *gorm.DB
}

// NewHubClientRepository creates a new instance of HubClientRepository.
func NewHubClientRepository(db *gorm.DB) HubClientRepository {
	return &hubClientRepository{
		base: NewBaseRepository[*models.HubClient](db),
		db:   db,
	}
}

// Pagination retrieves paginated hub clients from the database.
func (r *hubClientRepository) Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.HubClient, int64, error) {
	var clients []models.HubClient
	var total int64

	query := r.db.Model(&models.HubClient{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if active != nil {
		query = query.Where("active = ?", *active)
	}

	query.Count(&total)

	if sortField != "" {
		if sortOrder != "desc" {
			sortOrder = "asc"
		}
		query = query.Order(sortField + " " + sortOrder)
	}

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&clients).Error

	return clients, total, err
}

// GetAll retrieves all hub clients from the database with filtering and sorting.
func (r *hubClientRepository) GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error) {
	var clients []models.HubClient

	query := r.db.Model(&models.HubClient{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	if active != nil {
		query = query.Where("active = ?", *active)
	}

	if sortField != "" {
		if sortOrder != "desc" {
			sortOrder = "asc"
		}
		query = query.Order(sortField + " " + sortOrder)
	}

	err := query.Find(&clients).Error
	return clients, err
}

// GetByID retrieves a single hub client by ID using BaseRepository.
func (r *hubClientRepository) GetByID(id uint) (*models.HubClient, error) {
	return r.base.GetByID(id)
}

// Create inserts a new hub client into the database using BaseRepository.
func (r *hubClientRepository) Create(hubClient *models.HubClient) error {
	return r.base.Create(hubClient)
}

// Update modifies an existing hub client using BaseRepository.
func (r *hubClientRepository) Update(hubClient *models.HubClient) error {
	return r.base.Update(hubClient)
}

// Delete removes a hub client from the database by its ID using BaseRepository.
func (r *hubClientRepository) Delete(id uint) error {
	return r.base.Delete(id)
}

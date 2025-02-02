package repositories

import (
	"go-modules-api/internal/models"
	"gorm.io/gorm"
)

// RoleRepository defines the interface for database operations related to roles.
type RoleRepository interface {
	BaseRepositoryInterface[*models.Role]
	Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.Role, int64, error)
	GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.Role, error)
}

type roleRepository struct {
	base *BaseRepository[*models.Role]
	db   *gorm.DB
}

// NewRoleRepository creates a new instance of RoleRepository.
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		base: NewBaseRepository[*models.Role](db),
		db:   db,
	}
}

// Pagination returns paginated roles based on search criteria, active status, sorting, and pagination parameters.
func (r *roleRepository) Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.Role, int64, error) {
	var roles []models.Role
	var total int64

	query := r.db.Model(&models.Role{}).Scopes(scopeNotDeleted)

	if search != "" {
		// Search by name or slug using a case-insensitive LIKE query.
		query = query.Where("name ILIKE ? OR slug ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if active != nil {
		query = query.Where("active = ?", *active)
	}

	// Count the total records after applying filters.
	query.Count(&total)

	if sortField != "" {
		if sortOrder != "desc" {
			sortOrder = "asc"
		}
		query = query.Order(sortField + " " + sortOrder)
	}

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&roles).Error

	return roles, total, err
}

// GetAll returns all roles based on search criteria and sorting options.
func (r *roleRepository) GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.Role, error) {
	var roles []models.Role

	query := r.db.Model(&models.Role{}).Scopes(scopeNotDeleted)

	if search != "" {
		query = query.Where("name ILIKE ? OR slug ILIKE ?", "%"+search+"%", "%"+search+"%")
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

	err := query.Find(&roles).Error
	return roles, err
}

// GetByID returns a single role by its ID using BaseRepository.
func (r *roleRepository) GetByID(id uint) (*models.Role, error) {
	return r.base.GetByID(id)
}

// Create inserts a new role into the database using BaseRepository.
func (r *roleRepository) Create(role *models.Role) error {
	return r.base.Create(role)
}

// Update modifies an existing role using BaseRepository.
func (r *roleRepository) Update(role *models.Role) error {
	return r.base.Update(role)
}

// Delete removes a role from the database by its ID using BaseRepository.
func (r *roleRepository) Delete(id uint) error {
	return r.base.Delete(id)
}

// SoftDelete marks a role as deleted without actually removing it from the database using BaseRepository.
func (r *roleRepository) SoftDelete(role *models.Role) error {
	return r.base.SoftDelete(role)
}

package repositories

import (
	"time"

	"go-modules-api/internal/models"

	"gorm.io/gorm"
)

// BaseRepositoryInterface defines common CRUD operations for any model type that has an ID.
type BaseRepositoryInterface[T any] interface {
	GetByID(id uint) (T, error)
	Create(entity T) error
	Update(entity T) error
	Delete(id uint) error
	SoftDelete(entity T) error
}

// BaseRepository provides common CRUD operations for any model type that has an ID.
type BaseRepository[T models.IDGetter] struct {
	db *gorm.DB
}

// NewBaseRepository creates a new BaseRepository for the given model type.
func NewBaseRepository[T models.IDGetter](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// scopeNotDeleted adds a condition to ignore records where IsDeleted is true.
func scopeNotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("is_deleted = ?", false)
}

// GetByID retrieves a record by its ID.
func (r *BaseRepository[T]) GetByID(id uint) (T, error) {
	var entity T
	if err := r.db.Scopes(scopeNotDeleted).First(&entity, id).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

// Create inserts a new record into the database.
func (r *BaseRepository[T]) Create(entity T) error {
	return r.db.Create(entity).Error
}

// Update updates an existing record.
func (r *BaseRepository[T]) Update(entity T) error {
	return r.db.Scopes(scopeNotDeleted).Where("id = ?", entity.GetID()).Updates(entity).Error
}

// Delete removes a record from the database by its ID.
func (r *BaseRepository[T]) Delete(id uint) error {
	var entity T
	return r.db.Scopes(scopeNotDeleted).Delete(&entity, id).Error
}

// SoftDelete marks a record as deleted without actually removing it from the database.
func (r *BaseRepository[T]) SoftDelete(entity T) error {
	return r.db.Model(entity).Updates(map[string]interface{}{"is_deleted": true, "deleted_at": time.Now()}).Error
}

// Ensure BaseRepository implements BaseRepositoryInterface.
var _ BaseRepositoryInterface[models.IDGetter] = &BaseRepository[models.IDGetter]{}

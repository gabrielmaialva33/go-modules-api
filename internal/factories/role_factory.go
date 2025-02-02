package factories

import (
	"go-modules-api/internal/models"

	"github.com/brianvoe/gofakeit/v6"
)

// RoleFactory creates a fake Role instance
func RoleFactory() *models.Role {
	return &models.Role{
		Name: gofakeit.Name(),
		Slug: gofakeit.Username(),
	}
}

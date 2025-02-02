package factories

import (
	"go-modules-api/internal/models"

	"github.com/brianvoe/gofakeit/v6"
)

// HubClientFactory creates a fake HubClient instance
func HubClientFactory() *models.HubClient {
	return &models.HubClient{
		Name:       gofakeit.Name(),
		ExternalID: gofakeit.UUID(),
	}
}

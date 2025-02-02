package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go-modules-api/internal/dto"
	"go-modules-api/internal/models"
)

// MockHubClientService is a mock implementation of the HubClientService interface
type MockHubClientService struct {
	mock.Mock
}

func (m *MockHubClientService) PaginateHubClients(params dto.PaginatedHubClientDTO) ([]models.HubClient, int64, error) {
	args := m.Called(params)
	return args.Get(0).([]models.HubClient), args.Get(1).(int64), args.Error(2)
}

func (m *MockHubClientService) ListHubClients(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error) {
	args := m.Called(search, active, sortField, sortOrder)
	return args.Get(0).([]models.HubClient), args.Error(1)
}

func (m *MockHubClientService) GetHubClientByID(id uint) (*models.HubClient, error) {
	args := m.Called(id)
	return args.Get(0).(*models.HubClient), args.Error(1)
}

func (m *MockHubClientService) CreateHubClient(hubClient *models.HubClient) error {
	args := m.Called(hubClient)
	return args.Error(0)
}

func (m *MockHubClientService) UpdateHubClient(hubClient *models.HubClient) error {
	args := m.Called(hubClient)
	return args.Error(0)
}

func (m *MockHubClientService) DeleteHubClient(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockHubClientService) SoftDeleteHubClient(hubClient *models.HubClient) error {
	args := m.Called(hubClient)
	return args.Error(0)
}

func TestHubClientHandler_PaginateHubClients(t *testing.T) {
	mockService := new(MockHubClientService)
	handler := NewHubClientHandler(mockService)

	app := fiber.New()
	app.Get("/hub_clients/paginate", handler.PaginateHubClients)

	params := dto.PaginatedHubClientDTO{
		Search:    "Client",
		SortField: "id",
		SortOrder: "asc",
		Page:      1,
		PageSize:  10,
	}
	mockClients := []models.HubClient{{BaseID: models.BaseID{ID: 1}, Name: "Client1"}, {BaseID: models.BaseID{ID: 2}, Name: "Client2"}}
	mockTotal := int64(2)

	mockService.On("PaginateHubClients", params).Return(mockClients, mockTotal, nil)

	req := httptest.NewRequest("GET", "/hub_clients/paginate?search=Client&sort_field=id&sort_order=asc&page=1&page_size=10", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHubClientHandler_ListHubClients(t *testing.T) {
	mockService := new(MockHubClientService)
	handler := NewHubClientHandler(mockService)

	app := fiber.New()
	app.Get("/hub_clients", handler.ListHubClients)

	search := "Client"
	active := true
	sortField := "id"
	sortOrder := "asc"
	mockClients := []models.HubClient{{BaseID: models.BaseID{ID: 1}, Name: "Client1"}}

	mockService.On("ListHubClients", search, &active, sortField, sortOrder).Return(mockClients, nil)

	req := httptest.NewRequest("GET", "/hub_clients?search=Client&active=true&sort_field=id&sort_order=asc", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHubClientHandler_GetHubClientByID(t *testing.T) {
	mockService := new(MockHubClientService)
	handler := NewHubClientHandler(mockService)

	app := fiber.New()
	app.Get("/hub_clients/:id", handler.GetHubClientByID)

	mockID := uint(1)
	mockClient := &models.HubClient{BaseID: models.BaseID{ID: mockID}, Name: "Client1"}

	mockService.On("GetHubClientByID", mockID).Return(mockClient, nil)

	req := httptest.NewRequest("GET", "/hub_clients/1", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHubClientHandler_CreateHubClient(t *testing.T) {
	mockService := new(MockHubClientService)
	handler := NewHubClientHandler(mockService)

	app := fiber.New()
	app.Post("/hub_clients", handler.CreateHubClient)

	mockClient := &models.HubClient{Name: "Client1", ExternalID: "1"}
	mockService.On("CreateHubClient", mockClient).Return(nil)

	payload := `{"name":"Client1","external_id": "1"}`
	req := httptest.NewRequest("POST", "/hub_clients", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestHubClientHandler_UpdateHubClient(t *testing.T) {
	mockService := new(MockHubClientService)
	handler := NewHubClientHandler(mockService)

	app := fiber.New()
	app.Put("/hub_clients/:id", handler.UpdateHubClient)

	mockID := uint(1)
	mockClient := &models.HubClient{BaseID: models.BaseID{ID: mockID}, Name: "Client1", ExternalID: "1"}
	mockService.On("GetHubClientByID", mockID).Return(mockClient, nil)
	updatedClient := &models.HubClient{BaseID: models.BaseID{ID: mockID}, Name: "Client1", ExternalID: "1"}
	mockService.On("UpdateHubClient", updatedClient).Return(nil)

	payload := `{"name":"Client1","external_id": "1"}`
	req := httptest.NewRequest("PUT", "/hub_clients/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestHubClientHandler_DeleteHubClient(t *testing.T) {
	mockService := new(MockHubClientService)
	handler := NewHubClientHandler(mockService)

	app := fiber.New()
	app.Delete("/hub_clients/:id", handler.SoftDeleteHubClient)

	mockID := uint(1)
	mockClient := &models.HubClient{BaseID: models.BaseID{ID: mockID}, Name: "Client1"}
	mockService.On("GetHubClientByID", mockID).Return(mockClient, nil)
	mockService.On("SoftDeleteHubClient", mockClient).Return(nil)

	req := httptest.NewRequest("DELETE", "/hub_clients/1", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

	mockService.AssertExpectations(t)
}

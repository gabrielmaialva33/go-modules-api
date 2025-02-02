package services

import (
	"go-modules-api/internal/dto"
	"go-modules-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHubClientRepository is a mock implementation of the HubClientRepository interface
type MockHubClientRepository struct {
	mock.Mock
}

func (m *MockHubClientRepository) Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.HubClient, int64, error) {
	args := m.Called(search, active, sortField, sortOrder, page, pageSize)
	return args.Get(0).([]models.HubClient), args.Get(1).(int64), args.Error(2)
}

func (m *MockHubClientRepository) GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.HubClient, error) {
	args := m.Called(search, active, sortField, sortOrder)
	return args.Get(0).([]models.HubClient), args.Error(1)
}

func (m *MockHubClientRepository) GetByID(id uint) (*models.HubClient, error) {
	args := m.Called(id)
	return args.Get(0).(*models.HubClient), args.Error(1)
}

func (m *MockHubClientRepository) Create(hubClient *models.HubClient) error {
	args := m.Called(hubClient)
	return args.Error(0)
}

func (m *MockHubClientRepository) Update(hubClient *models.HubClient) error {
	args := m.Called(hubClient)
	return args.Error(0)
}

func (m *MockHubClientRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestHubClientService_PaginateHubClients(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := NewHubClientService(mockRepo)

	mockResults := []models.HubClient{{ID: 1, Name: "Client1"}, {ID: 2, Name: "Client2"}}
	mockTotal := int64(2)
	mockParams := dto.PaginatedHubClientDTO{
		Page:      1,
		PageSize:  10,
		Search:    "Client",
		Active:    nil,
		SortField: "id",
		SortOrder: "asc",
	}

	mockRepo.On("Pagination", mockParams.Search, mockParams.Active, mockParams.SortField, mockParams.SortOrder, mockParams.Page, mockParams.PageSize).Return(mockResults, mockTotal, nil)

	result, total, err := service.PaginateHubClients(mockParams)

	assert.NoError(t, err)
	assert.Equal(t, mockTotal, total)
	assert.Equal(t, mockResults, result)

	mockRepo.AssertExpectations(t)
}

func TestHubClientService_ListHubClients(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := NewHubClientService(mockRepo)

	search := "Client"
	active := true
	sortField := "name"
	sortOrder := "asc"
	mockResults := []models.HubClient{{ID: 1, Name: "Client1"}}

	mockRepo.On("GetAll", search, &active, sortField, sortOrder).Return(mockResults, nil)

	result, err := service.ListHubClients(search, &active, sortField, sortOrder)

	assert.NoError(t, err)
	assert.Equal(t, mockResults, result)

	mockRepo.AssertExpectations(t)
}

func TestHubClientService_GetHubClientByID(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := NewHubClientService(mockRepo)

	mockID := uint(1)
	mockResult := &models.HubClient{ID: mockID, Name: "Client1"}

	mockRepo.On("GetByID", mockID).Return(mockResult, nil)

	result, err := service.GetHubClientByID(mockID)

	assert.NoError(t, err)
	assert.Equal(t, mockResult, result)

	mockRepo.AssertExpectations(t)
}

func TestHubClientService_CreateHubClient(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := NewHubClientService(mockRepo)

	mockClient := &models.HubClient{ID: 1, Name: "Client1"}

	mockRepo.On("Create", mockClient).Return(nil)

	err := service.CreateHubClient(mockClient)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestHubClientService_UpdateHubClient(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := NewHubClientService(mockRepo)

	mockClient := &models.HubClient{ID: 1, Name: "UpdatedClient"}

	mockRepo.On("Update", mockClient).Return(nil)

	err := service.UpdateHubClient(mockClient)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestHubClientService_DeleteHubClient(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := NewHubClientService(mockRepo)

	mockID := uint(1)

	mockRepo.On("Delete", mockID).Return(nil)

	err := service.DeleteHubClient(mockID)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

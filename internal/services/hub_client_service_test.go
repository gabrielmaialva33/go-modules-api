package services_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go-modules-api/internal/dto"
	"go-modules-api/internal/models"
	"go-modules-api/internal/services"
)

// ---------------------------
// MockHubClientRepository
// ---------------------------

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
	if args.Get(0) != nil {
		return args.Get(0).(*models.HubClient), args.Error(1)
	}
	return nil, args.Error(1)
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

// ---------------------------
// Service Test
// ---------------------------

func TestPaginateHubClients_Success(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	params := dto.PaginatedHubClientDTO{
		Search:    "test",
		Active:    nil,
		SortField: "name",
		SortOrder: "asc",
		Page:      1,
		PageSize:  10,
	}

	expectedClients := []models.HubClient{
		{Name: "Client 1"},
		{Name: "Client 2"},
	}
	expectedTotal := int64(2)

	mockRepo.
		On("Pagination", params.Search, params.Active, params.SortField, params.SortOrder, params.Page, params.PageSize).
		Return(expectedClients, expectedTotal, nil)

	clients, total, err := service.PaginateHubClients(params)
	assert.NoError(t, err)
	assert.Equal(t, expectedClients, clients)
	assert.Equal(t, expectedTotal, total)

	mockRepo.AssertExpectations(t)
}

func TestPaginateHubClients_Error(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	params := dto.PaginatedHubClientDTO{
		Search:    "",
		Active:    nil,
		SortField: "name",
		SortOrder: "asc",
		Page:      1,
		PageSize:  10,
	}

	expectedError := errors.New("db error")
	mockRepo.
		On("Pagination", params.Search, params.Active, params.SortField, params.SortOrder, params.Page, params.PageSize).
		Return([]models.HubClient{}, int64(0), expectedError)

	_, _, err := service.PaginateHubClients(params)
	assert.Error(t, err)
	// Compara a mensagem completa retornada por HandleDBError.
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestListHubClients_Success(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	search := "test"
	active := new(bool)
	*active = true
	sortField := "name"
	sortOrder := "asc"

	expectedClients := []models.HubClient{
		{Name: "Client 1"},
	}
	mockRepo.
		On("GetAll", search, active, sortField, sortOrder).
		Return(expectedClients, nil)

	clients, err := service.ListHubClients(search, active, sortField, sortOrder)
	assert.NoError(t, err)
	assert.Equal(t, expectedClients, clients)

	mockRepo.AssertExpectations(t)
}

func TestListHubClients_Error(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	search := ""
	var active *bool = nil
	sortField := "name"
	sortOrder := "asc"

	expectedError := errors.New("db error")
	mockRepo.
		On("GetAll", search, active, sortField, sortOrder).
		Return([]models.HubClient{}, expectedError)

	_, err := service.ListHubClients(search, active, sortField, sortOrder)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetHubClientByID_Success(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	clientID := uint(1)
	expectedClient := &models.HubClient{Name: "Client 1"}
	mockRepo.
		On("GetByID", clientID).
		Return(expectedClient, nil)

	client, err := service.GetHubClientByID(clientID)
	assert.NoError(t, err)
	assert.Equal(t, expectedClient, client)

	mockRepo.AssertExpectations(t)
}

func TestGetHubClientByID_Error(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	clientID := uint(1)
	expectedError := errors.New("not found")
	mockRepo.
		On("GetByID", clientID).
		Return(nil, expectedError)

	client, err := service.GetHubClientByID(clientID)
	assert.Error(t, err)
	assert.Nil(t, client)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCreateHubClient_Success(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	newClient := &models.HubClient{Name: "New Client"}
	mockRepo.
		On("Create", newClient).
		Return(nil)

	err := service.CreateHubClient(newClient)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateHubClient_Error(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	newClient := &models.HubClient{Name: "New Client"}
	expectedError := errors.New("create error")
	mockRepo.
		On("Create", newClient).
		Return(expectedError)

	err := service.CreateHubClient(newClient)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUpdateHubClient_Success(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	clientToUpdate := &models.HubClient{Name: "Updated Client"}
	mockRepo.
		On("Update", clientToUpdate).
		Return(nil)

	err := service.UpdateHubClient(clientToUpdate)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateHubClient_Error(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	clientToUpdate := &models.HubClient{Name: "Updated Client"}
	expectedError := errors.New("update error")
	mockRepo.
		On("Update", clientToUpdate).
		Return(expectedError)

	err := service.UpdateHubClient(clientToUpdate)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDeleteHubClient_Success(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	clientID := uint(1)
	mockRepo.
		On("Delete", clientID).
		Return(nil)

	err := service.DeleteHubClient(clientID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteHubClient_Error(t *testing.T) {
	mockRepo := new(MockHubClientRepository)
	service := services.NewHubClientService(mockRepo)

	clientID := uint(1)
	expectedError := errors.New("delete error")
	mockRepo.
		On("Delete", clientID).
		Return(expectedError)

	err := service.DeleteHubClient(clientID)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

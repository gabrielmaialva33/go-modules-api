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

// MockRoleService is a mock implementation of the RoleService interface.
type MockRoleService struct {
	mock.Mock
}

func (m *MockRoleService) PaginateRoles(params dto.PaginatedRoleDTO) ([]models.Role, int64, error) {
	args := m.Called(params)
	return args.Get(0).([]models.Role), args.Get(1).(int64), args.Error(2)
}

func (m *MockRoleService) ListRoles(search string, active *bool, sortField, sortOrder string) ([]models.Role, error) {
	args := m.Called(search, active, sortField, sortOrder)
	return args.Get(0).([]models.Role), args.Error(1)
}

func (m *MockRoleService) GetRoleByID(id uint) (*models.Role, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *MockRoleService) CreateRole(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

func (m *MockRoleService) UpdateRole(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

func (m *MockRoleService) DeleteRole(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoleService) SoftDeleteRole(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

func TestRoleHandler_PaginateRoles(t *testing.T) {
	mockService := new(MockRoleService)
	handler := NewRoleHandler(mockService)

	app := fiber.New()
	app.Get("/roles/paginate", handler.PaginateRoles)

	// Update the expected arguments to match defaults:
	expectedParams := dto.PaginatedRoleDTO{
		Page:      1,  // Handler defaults
		PageSize:  10, // Handler defaults
		Search:    "",
		Active:    nil,
		SortField: "id",  // Handler defaults
		SortOrder: "asc", // Handler defaults
	}
	// Mock return data
	mockRoles := []models.Role{
		{BaseID: models.BaseID{ID: 1}, Name: "Admin"},
		{BaseID: models.BaseID{ID: 2}, Name: "Staff"},
	}
	mockTotal := int64(2)

	// The mock now matches the actual parameters
	mockService.On("PaginateRoles", expectedParams).Return(mockRoles, mockTotal, nil)

	req := httptest.NewRequest("GET", "/roles/paginate", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestRoleHandler_ListRoles(t *testing.T) {
	mockService := new(MockRoleService)
	handler := NewRoleHandler(mockService)

	app := fiber.New()
	app.Get("/roles", handler.ListRoles)

	search := "admin"
	active := true
	sortField := "id"
	sortOrder := "asc"
	mockRoles := []models.Role{
		{BaseID: models.BaseID{ID: 1}, Name: "Admin"},
	}

	mockService.On("ListRoles", search, &active, sortField, sortOrder).Return(mockRoles, nil)

	req := httptest.NewRequest("GET", "/roles?search=admin&active=true&sort_field=id&sort_order=asc", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestRoleHandler_GetRoleByID(t *testing.T) {
	mockService := new(MockRoleService)
	handler := NewRoleHandler(mockService)

	app := fiber.New()
	app.Get("/roles/:id", handler.GetRoleByID)

	mockID := uint(1)
	mockRole := &models.Role{BaseID: models.BaseID{ID: mockID}, Name: "Admin"}

	mockService.On("GetRoleByID", mockID).Return(mockRole, nil)

	req := httptest.NewRequest("GET", "/roles/1", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestRoleHandler_CreateRole(t *testing.T) {
	mockService := new(MockRoleService)
	handler := NewRoleHandler(mockService)

	app := fiber.New()
	app.Post("/roles", handler.CreateRole)

	role := &models.Role{
		Name: "Admin",
		Slug: "admin",
		BaseAttributes: models.BaseAttributes{
			Active: true, // Matches RoleHandler default
		},
	}

	// Make sure the mock expects the 'Active' field to be true
	mockService.On("CreateRole", role).Return(nil)

	payload := `{"name":"Admin","slug":"admin"}`
	req := httptest.NewRequest("POST", "/roles", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}

func TestRoleHandler_UpdateRole(t *testing.T) {
	mockService := new(MockRoleService)
	handler := NewRoleHandler(mockService)

	app := fiber.New()
	app.Put("/roles/:id", handler.UpdateRole)

	mockID := uint(1)
	existingRole := &models.Role{BaseID: models.BaseID{ID: mockID}, Name: "Admin", Slug: "admin"}
	updatedRole := &models.Role{BaseID: models.BaseID{ID: mockID}, Name: "Admin", Slug: "admin"}

	// Mock existing role retrieval
	mockService.On("GetRoleByID", mockID).Return(existingRole, nil)
	// Mock update call
	mockService.On("UpdateRole", updatedRole).Return(nil)

	payload := `{"name":"Admin","slug":"admin"}`
	req := httptest.NewRequest("PUT", "/roles/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestRoleHandler_SoftDeleteRole(t *testing.T) {
	mockService := new(MockRoleService)
	handler := NewRoleHandler(mockService)

	app := fiber.New()
	app.Delete("/roles/:id", handler.SoftDeleteRole)

	mockID := uint(1)
	existingRole := &models.Role{BaseID: models.BaseID{ID: mockID}, Name: "Admin"}

	// Mock existing role retrieval
	mockService.On("GetRoleByID", mockID).Return(existingRole, nil)
	// Mock soft-delete call
	mockService.On("SoftDeleteRole", existingRole).Return(nil)

	req := httptest.NewRequest("DELETE", "/roles/1", nil)
	resp, err := app.Test(req, -1)
	require.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)

	mockService.AssertExpectations(t)
}

package services_test

import (
	"errors"
	"testing"

	"go-modules-api/internal/dto"
	"go-modules-api/internal/models"
	"go-modules-api/internal/services"
	"go-modules-api/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ---------------------------
// MockRoleRepository
// ---------------------------

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) Pagination(search string, active *bool, sortField string, sortOrder string, page int, pageSize int) ([]models.Role, int64, error) {
	args := m.Called(search, active, sortField, sortOrder, page, pageSize)
	return args.Get(0).([]models.Role), args.Get(1).(int64), args.Error(2)
}

func (m *MockRoleRepository) GetAll(search string, active *bool, sortField string, sortOrder string) ([]models.Role, error) {
	args := m.Called(search, active, sortField, sortOrder)
	return args.Get(0).([]models.Role), args.Error(1)
}

func (m *MockRoleRepository) GetByID(id uint) (*models.Role, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Role), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockRoleRepository) Create(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

func (m *MockRoleRepository) Update(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

func (m *MockRoleRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockRoleRepository) SoftDelete(role *models.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

// ---------------------------
// Service Test
// ---------------------------

func TestPaginateRoles_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	params := dto.PaginatedRoleDTO{
		Search:    "test",
		Active:    nil,
		SortField: "name",
		SortOrder: "asc",
		Page:      1,
		PageSize:  10,
	}

	expectedRoles := []models.Role{
		{Name: "Role 1"},
		{Name: "Role 2"},
	}
	expectedTotal := int64(2)

	mockRepo.
		On("Pagination", params.Search, params.Active, params.SortField, params.SortOrder, params.Page, params.PageSize).
		Return(expectedRoles, expectedTotal, nil)

	roles, total, err := service.PaginateRoles(params)
	assert.NoError(t, err)
	assert.Equal(t, expectedRoles, roles)
	assert.Equal(t, expectedTotal, total)

	mockRepo.AssertExpectations(t)
}

func TestPaginateRoles_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	params := dto.PaginatedRoleDTO{
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
		Return([]models.Role{}, int64(0), expectedError)

	_, _, err := service.PaginateRoles(params)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestListRoles_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	search := "test"
	active := utils.BoolPtr(true)
	sortField := "name"
	sortOrder := "asc"

	expectedRoles := []models.Role{
		{Name: "Role 1"},
	}
	mockRepo.
		On("GetAll", search, active, sortField, sortOrder).
		Return(expectedRoles, nil)

	roles, err := service.ListRoles(search, active, sortField, sortOrder)
	assert.NoError(t, err)
	assert.Equal(t, expectedRoles, roles)

	mockRepo.AssertExpectations(t)
}

func TestListRoles_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	search := ""
	var active *bool = nil
	sortField := "name"
	sortOrder := "asc"

	expectedError := errors.New("db error")
	mockRepo.
		On("GetAll", search, active, sortField, sortOrder).
		Return([]models.Role{}, expectedError)

	_, err := service.ListRoles(search, active, sortField, sortOrder)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetRoleByID_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	roleID := uint(1)
	expectedRole := &models.Role{Name: "Role 1"}
	mockRepo.
		On("GetByID", roleID).
		Return(expectedRole, nil)

	role, err := service.GetRoleByID(roleID)
	assert.NoError(t, err)
	assert.Equal(t, expectedRole, role)

	mockRepo.AssertExpectations(t)
}

func TestGetRoleByID_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	roleID := uint(1)
	expectedError := errors.New("not found")
	mockRepo.
		On("GetByID", roleID).
		Return(nil, expectedError)

	role, err := service.GetRoleByID(roleID)
	assert.Error(t, err)
	assert.Nil(t, role)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCreateRole_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	newRole := &models.Role{Name: "New Role"}
	mockRepo.
		On("Create", newRole).
		Return(nil)

	err := service.CreateRole(newRole)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestCreateRole_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	newRole := &models.Role{Name: "New Role"}
	expectedError := errors.New("create error")
	mockRepo.
		On("Create", newRole).
		Return(expectedError)

	err := service.CreateRole(newRole)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUpdateRole_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	roleToUpdate := &models.Role{Name: "Updated Role"}
	mockRepo.
		On("Update", roleToUpdate).
		Return(nil)

	err := service.UpdateRole(roleToUpdate)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUpdateRole_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	roleToUpdate := &models.Role{Name: "Updated Role"}
	expectedError := errors.New("update error")
	mockRepo.
		On("Update", roleToUpdate).
		Return(expectedError)

	err := service.UpdateRole(roleToUpdate)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDeleteRole_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	roleID := uint(1)
	mockRepo.
		On("Delete", roleID).
		Return(nil)

	err := service.DeleteRole(roleID)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteRole_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	roleID := uint(1)
	expectedError := errors.New("delete error")
	mockRepo.
		On("Delete", roleID).
		Return(expectedError)

	err := service.DeleteRole(roleID)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestSoftDeleteRole_Success(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	role := &models.Role{Name: "Role to soft delete"}
	mockRepo.
		On("SoftDelete", role).
		Return(nil)

	err := service.SoftDeleteRole(role)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestSoftDeleteRole_Error(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	service := services.NewRoleService(mockRepo)

	role := &models.Role{Name: "Role to soft delete"}
	expectedError := errors.New("soft delete error")
	mockRepo.
		On("SoftDelete", role).
		Return(expectedError)

	err := service.SoftDeleteRole(role)
	assert.Error(t, err)
	assert.Equal(t, "[500] internal_server_error: A database error occurred", err.Error())

	mockRepo.AssertExpectations(t)
}

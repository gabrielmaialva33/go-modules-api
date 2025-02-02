package services

import (
	"go-modules-api/internal/dto"
	"go-modules-api/internal/models"
	"go-modules-api/internal/repositories"
	"go-modules-api/utils"
)

type RoleService interface {
	PaginateRoles(params dto.PaginatedRoleDTO) ([]models.Role, int64, error)
	ListRoles(search string, active *bool, sortField string, sortOrder string) ([]models.Role, error)
	GetRoleByID(id uint) (*models.Role, error)
	CreateRole(role *models.Role) error
	UpdateRole(role *models.Role) error
	DeleteRole(id uint) error
}

type roleService struct {
	repo repositories.RoleRepository
}

func NewRoleService(repo repositories.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) PaginateRoles(params dto.PaginatedRoleDTO) ([]models.Role, int64, error) {
	roles, total, err := s.repo.Pagination(
		params.Search,
		params.Active,
		params.SortField,
		params.SortOrder,
		params.Page,
		params.PageSize,
	)
	return roles, total, utils.HandleDBError(err)
}

func (s *roleService) ListRoles(search string, active *bool, sortField string, sortOrder string) ([]models.Role, error) {
	roles, err := s.repo.GetAll(search, active, sortField, sortOrder)
	return roles, utils.HandleDBError(err)
}

func (s *roleService) GetRoleByID(id uint) (*models.Role, error) {
	role, err := s.repo.GetByID(id)
	return role, utils.HandleDBError(err)
}

func (s *roleService) CreateRole(role *models.Role) error {
	return utils.HandleDBError(s.repo.Create(role))
}

func (s *roleService) UpdateRole(role *models.Role) error {
	return utils.HandleDBError(s.repo.Update(role))
}

func (s *roleService) DeleteRole(id uint) error {
	return utils.HandleDBError(s.repo.Delete(id))
}

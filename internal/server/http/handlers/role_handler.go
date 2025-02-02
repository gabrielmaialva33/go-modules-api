package handlers

import (
	"errors"
	"strconv"
	"strings"

	"go-modules-api/internal/dto"
	"go-modules-api/internal/exceptions"
	"go-modules-api/internal/models"
	"go-modules-api/internal/services"
	"go-modules-api/utils"

	"github.com/gofiber/fiber/v2"
)

// RoleHandler handles HTTP requests for roles
type RoleHandler struct {
	service services.RoleService
}

// NewRoleHandler creates a new RoleHandler
func NewRoleHandler(service services.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

// PaginateRoles handles GET /roles/paginate
func (h *RoleHandler) PaginateRoles(c *fiber.Ctx) error {
	params := dto.PaginatedRoleDTO{
		Search:    c.Query("search", ""),
		SortField: c.Query("sort_field", "id"),
		SortOrder: strings.ToLower(c.Query("sort_order", "asc")),
		Page:      c.QueryInt("page", 1),
		PageSize:  c.QueryInt("page_size", 10),
	}

	activeStr := c.Query("active")
	if activeStr != "" {
		activeBool, err := strconv.ParseBool(activeStr)
		if err == nil {
			params.Active = &activeBool
		}
	}

	validationErrors := utils.ValidateStruct(params)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}

	roles, total, err := h.service.PaginateRoles(params)
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	meta := utils.GeneratePaginationMeta(total, params.Page, params.PageSize)

	return c.JSON(fiber.Map{"data": roles, "meta": meta})
}

// ListRoles handles GET /roles with filtering and sorting
func (h *RoleHandler) ListRoles(c *fiber.Ctx) error {
	params := dto.ListRoleDTO{
		Search:    c.Query("search", ""),
		SortField: c.Query("sort_field", "id"),
		SortOrder: strings.ToLower(c.Query("sort_order", "asc")),
	}

	activeStr := c.Query("active")
	if activeStr != "" {
		activeBool, err := strconv.ParseBool(activeStr)
		if err == nil {
			params.Active = &activeBool
		}
	}

	validationErrors := utils.ValidateStruct(params)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}

	roles, err := h.service.ListRoles(params.Search, params.Active, params.SortField, params.SortOrder)
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.JSON(roles)
}

// GetRoleByID handles GET /roles/:id
func (h *RoleHandler) GetRoleByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	role, err := h.service.GetRoleByID(uint(id))
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.JSON(role)
}

// CreateRole handles POST /roles
func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	var payload dto.CreateRoleDTO

	// Parse JSON body
	if err := c.BodyParser(&payload); err != nil {
		return exceptions.BadRequest("Invalid request body", nil).Response(c)
	}

	// Validate input
	validationErrors := utils.ValidateStruct(payload)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}

	// Create Role model from DTO
	role := &models.Role{
		BaseID: models.BaseID{},
		Name:   payload.Name,
		Slug:   payload.Slug,
		BaseAttributes: models.BaseAttributes{
			Active:    true,
			IsDeleted: false,
		},
		BaseTimestamps: models.BaseTimestamps{},
	}

	// Call service
	err := h.service.CreateRole(role)
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.Status(fiber.StatusCreated).JSON(role)
}

// UpdateRole handles PUT /roles/:id
func (h *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	var payload dto.UpdateRoleDTO
	if err := c.BodyParser(&payload); err != nil {
		return exceptions.BadRequest("Invalid request body", nil).Response(c)
	}

	// Validate input
	validationErrors := utils.ValidateStruct(payload)
	if len(validationErrors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationErrors,
		})
	}

	_, err = h.service.GetRoleByID(uint(id))
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	// Update entity with new values
	role := &models.Role{
		BaseID: models.BaseID{ID: uint(id)},
		Name:   payload.Name,
		Slug:   payload.Slug,
	}

	// Call service
	if err := h.service.UpdateRole(role); err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.JSON(role)
}

// SoftDeleteRole handles DELETE /roles/:id
func (h *RoleHandler) SoftDeleteRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	role, err := h.service.GetRoleByID(uint(id))
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	if err := h.service.SoftDeleteRole(role); err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

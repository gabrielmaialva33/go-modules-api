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

// HubClientHandler handles HTTP requests for hub clients
type HubClientHandler struct {
	service services.HubClientService
}

func NewHubClientHandler(service services.HubClientService) *HubClientHandler {
	return &HubClientHandler{service: service}
}

// PaginateHubClients handles GET /hub_clients/paginate
func (h *HubClientHandler) PaginateHubClients(c *fiber.Ctx) error {
	params := dto.PaginatedHubClientDTO{
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

	clients, total, err := h.service.PaginateHubClients(params)
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	meta := utils.GeneratePaginationMeta(total, params.Page, params.PageSize)

	return c.JSON(fiber.Map{"data": clients, "meta": meta})
}

// ListHubClients handles GET /hub_clients with filtering and sorting
func (h *HubClientHandler) ListHubClients(c *fiber.Ctx) error {
	params := dto.ListHubClientDTO{
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

	clients, err := h.service.ListHubClients(params.Search, params.Active, params.SortField, params.SortOrder)
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.JSON(clients)
}

// GetHubClientByID handles GET /hub_clients/:id
func (h *HubClientHandler) GetHubClientByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	client, err := h.service.GetHubClientByID(uint(id))
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.JSON(client)
}

// CreateHubClient handles POST /hub_clients
func (h *HubClientHandler) CreateHubClient(c *fiber.Ctx) error {
	var payload dto.CreateHubClientDTO

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

	// Create HubClient model from DTO
	hubClient := &models.HubClient{
		Name:       payload.Name,
		ExternalID: payload.ExternalID,
	}

	// Call service
	err := h.service.CreateHubClient(hubClient)
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.Status(fiber.StatusCreated).JSON(hubClient)
}

// UpdateHubClient handles PUT /hub_clients/:id
func (h *HubClientHandler) UpdateHubClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	var payload dto.UpdateHubClientDTO
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

	_, err = h.service.GetHubClientByID(uint(id))
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	// Update entity with new values
	client := &models.HubClient{
		BaseID:     models.BaseID{ID: uint(id)},
		Name:       payload.Name,
		ExternalID: payload.ExternalID,
	}

	// Call service
	if err := h.service.UpdateHubClient(client); err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.JSON(client)
}

// SoftDeleteHubClient handles DELETE /hub_clients/:id
func (h *HubClientHandler) SoftDeleteHubClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	client, err := h.service.GetHubClientByID(uint(id))
	if err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	if err := h.service.SoftDeleteHubClient(client); err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

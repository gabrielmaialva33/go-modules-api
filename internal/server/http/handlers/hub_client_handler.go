package handlers

import (
	"errors"

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

// GetAllHubClients handles GET /hub_clients
func (h *HubClientHandler) GetAllHubClients(c *fiber.Ctx) error {
	clients, err := h.service.GetAllHubClients()
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

	var payload dto.CreateHubClientDTO
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

	// Update entity with new values
	client := &models.HubClient{
		ID:         uint(id),
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

// DeleteHubClient handles DELETE /hub_clients/:id
func (h *HubClientHandler) DeleteHubClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	if err := h.service.DeleteHubClient(uint(id)); err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

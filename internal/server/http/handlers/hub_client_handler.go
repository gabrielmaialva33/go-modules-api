package handlers

import (
	"errors"
	"go-modules-api/internal/exceptions"
	"go-modules-api/internal/models"
	"go-modules-api/internal/services"

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
		return err.(*exceptions.APIException).Response(c)
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
		return err.(*exceptions.APIException).Response(c)
	}
	return c.JSON(client)
}

// CreateHubClient handles POST /hub_clients
func (h *HubClientHandler) CreateHubClient(c *fiber.Ctx) error {
	var client models.HubClient
	if err := c.BodyParser(&client); err != nil {
		return exceptions.BadRequest("Invalid request body", nil).Response(c)
	}

	if err := h.service.CreateHubClient(&client); err != nil {
		var apiErr *exceptions.APIException
		if errors.As(err, &apiErr) {
			return apiErr.Response(c)
		}
		return exceptions.InternalServerError("An unexpected error occurred", nil).Response(c)
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}

// UpdateHubClient handles PUT /hub_clients/:id
func (h *HubClientHandler) UpdateHubClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return exceptions.BadRequest("Invalid ID format", fiber.Map{"field": "id", "value": c.Params("id")}).Response(c)
	}

	var client models.HubClient
	if err := c.BodyParser(&client); err != nil {
		return exceptions.BadRequest("Invalid request body", nil).Response(c)
	}

	client.ID = uint(id)
	if err := h.service.UpdateHubClient(&client); err != nil {
		return err.(*exceptions.APIException).Response(c)
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
		return err.(*exceptions.APIException).Response(c)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

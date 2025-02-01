package handlers

import (
	"go-modules-api/internal/models"
	"go-modules-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

// HubClientHandler handles HTTP requests for hub clients
type HubClientHandler struct {
	service services.HubClientService
}

// NewHubClientHandler creates a new handler instance
func NewHubClientHandler(service services.HubClientService) *HubClientHandler {
	return &HubClientHandler{service: service}
}

// GetAllHubClients handles GET /hub_clients
func (h *HubClientHandler) GetAllHubClients(c *fiber.Ctx) error {
	clients, err := h.service.GetAllHubClients()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch hub clients"})
	}
	return c.JSON(clients)
}

// GetHubClientByID handles GET /hub_clients/:id
func (h *HubClientHandler) GetHubClientByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	client, err := h.service.GetHubClientByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Hub client not found"})
	}

	return c.JSON(client)
}

// CreateHubClient handles POST /hub_clients
func (h *HubClientHandler) CreateHubClient(c *fiber.Ctx) error {
	var client models.HubClient
	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   "bad_request",
			"message": "Invalid request body",
		})
	}

	if err := h.service.CreateHubClient(&client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create hub client"})
	}

	return c.Status(fiber.StatusCreated).JSON(client)
}

// UpdateHubClient handles PUT /hub_clients/:id
func (h *HubClientHandler) UpdateHubClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var client models.HubClient
	if err := c.BodyParser(&client); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	client.ID = uint(id)
	if err := h.service.UpdateHubClient(&client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update hub client"})
	}

	return c.JSON(client)
}

// DeleteHubClient handles DELETE /hub_clients/:id
func (h *HubClientHandler) DeleteHubClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.service.DeleteHubClient(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete hub client"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

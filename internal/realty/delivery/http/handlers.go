package http

import (
	"github.com/gofiber/fiber/v3"
	"realty/internal/realty"
	"realty/pkg/logger"
	"strconv"
)

//go:generate ifacemaker -f handlers.go -o ../../handlers.go -i Handlers -s RealtyHandler -p realty -y "Controller describes methods, implemented by the http package."
type RealtyHandler struct {
	realtyUC realty.UseCase
	logger   *logger.ApiLogger
}

func NewRealtyHandler(realtyUC realty.UseCase, logger *logger.ApiLogger) *RealtyHandler {
	return &RealtyHandler{realtyUC: realtyUC, logger: logger}
}

func (h RealtyHandler) CreateHouse() fiber.Handler {
	return func(c fiber.Ctx) error {
		userType := c.Locals("userType").(string)
		if userType != "Moderator" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}
		var req realty.CreateHouseRequest
		if err := c.Bind().Body(&req); err != nil {
			h.logger.Errorf("Failed to parse request body", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		house, err := h.realtyUC.CreateHouse(&req)
		if err != nil {
			h.logger.Errorf("Failed to create house", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(house)
	}
}

func (h RealtyHandler) CreateFlat() fiber.Handler {
	return func(c fiber.Ctx) error {
		var req realty.CreateFlatRequest
		if err := c.Bind().Body(&req); err != nil {
			h.logger.Errorf("Failed to parse request body", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		flat, err := h.realtyUC.CreateFlat(&req)
		if err != nil {
			h.logger.Errorf("Failed to create flat", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(flat)
	}
}

func (h RealtyHandler) UpdateFlatStatus() fiber.Handler {
	return func(c fiber.Ctx) error {
		userType := c.Locals("userType").(string)
		if userType != "Moderator" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}
		var req realty.UpdateFlatStatusRequest
		if err := c.Bind().Body(&req); err != nil {
			h.logger.Errorf("Failed to parse request body", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		validStatuses := map[string]bool{"created": true, "approved": true, "declined": true, "on moderation": true}
		if !validStatuses[req.NewStatus] {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status"})
		}

		flat, err := h.realtyUC.GetFlatByID(req.FlatID)
		if err != nil {
			h.logger.Errorf("Failed to get flat by ID", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		if req.NewStatus == "on moderation" && flat.Status != "created" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Flat can only be put on moderation from created status"})
		}

		if (req.NewStatus == "approved" || req.NewStatus == "declined") && flat.Status != "on moderation" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Flat can only be approved or declined from on moderation status"})
		}

		updatedFlat, err := h.realtyUC.UpdateFlatStatus(&req)
		if err != nil {
			h.logger.Errorf("Failed to update flat status", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(updatedFlat)
	}
}

func (h RealtyHandler) GetFlatsByHouseID() fiber.Handler {
	return func(c fiber.Ctx) error {
		houseIDStr := c.Params("id")
		houseID, err := strconv.ParseInt(houseIDStr, 10, 64)
		userType := c.Locals("userType").(string)
		flats, err := h.realtyUC.GetFlatsByHouseID(houseID, userType)
		if err != nil {
			h.logger.Errorf("Failed to get flats by house ID", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(flats)
	}
}

func (h RealtyHandler) SubscribeToHouse() fiber.Handler {
	return func(c fiber.Ctx) error {
		houseIDStr := c.Params("id")
		houseID, err := strconv.ParseInt(houseIDStr, 10, 64)
		if err != nil {
			h.logger.Errorf("Invalid house ID", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid house ID"})
		}

		var req realty.SubscribeToHouseRequest
		if err = c.Bind().Body(&req); err != nil {
			h.logger.Errorf("Failed to parse request body", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		err = h.realtyUC.SubscribeToHouse(req.Email, houseID)
		if err != nil {
			h.logger.Errorf("Failed to subscribe to house", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Subscribed successfully"})
	}
}

package http

import (
	"github.com/gofiber/fiber/v3"
	"realty/internal/auth"
	"realty/pkg/logger"
)

//go:generate ifacemaker -f handlers.go -o ../../handlers.go -i Handlers -s AuthHandler -p auth -y "Controller describes methods, implemented by the http package."

type AuthHandler struct {
	authUC auth.UseCase
	logger *logger.ApiLogger
}

func NewAuthHandler(authUC auth.UseCase, logger *logger.ApiLogger) *AuthHandler {
	return &AuthHandler{authUC: authUC, logger: logger}
}

func (h AuthHandler) Register() fiber.Handler {
	return func(c fiber.Ctx) error {
		var req auth.RegisterRequest

		if err := c.Bind().Body(&req); err != nil {
			h.logger.Errorf("Failed to parse request body", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		res, err := h.authUC.Register(&req)
		if err != nil {
			h.logger.Errorf("Failed to register", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}

func (h AuthHandler) Login() fiber.Handler {
	return func(c fiber.Ctx) error {
		var req auth.LoginRequest

		if err := c.Bind().Body(&req); err != nil {
			h.logger.Errorf("Failed to parse request body", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		res, err := h.authUC.Login(&req)
		if err != nil {
			h.logger.Errorf("Failed to register", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "InternalServerError"})
		}

		return c.Status(fiber.StatusOK).JSON(res)
	}
}

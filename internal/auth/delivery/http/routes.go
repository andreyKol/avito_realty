package http

import (
	"github.com/gofiber/fiber/v3"
	"realty/internal/auth"
)

func MapAuthRoutes(r fiber.Router, h auth.Handlers) {
	r.Post(`/register`, h.Register())
	r.Post(`/login`, h.Login())
}

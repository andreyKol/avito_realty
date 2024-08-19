package http

import (
	"github.com/gofiber/fiber/v3"
	"realty/internal/realty"
)

func MapRealtyRoutes(r fiber.Router, h realty.Handlers) {
	r.Post(`/house/create`, h.CreateHouse())
	r.Post(`/flat/create`, h.CreateFlat())
	r.Patch(`/flat/update`, h.UpdateFlatStatus())
	r.Get("/house/:id", h.GetFlatsByHouseID())
	r.Post("/house/:id/subscribe", h.SubscribeToHouse())
}

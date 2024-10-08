// Code generated by ifacemaker; DO NOT EDIT.

package realty

import (
	"github.com/gofiber/fiber/v3"
)

// Controller describes methods, implemented by the http package.
type Handlers interface {
	CreateHouse() fiber.Handler
	CreateFlat() fiber.Handler
	UpdateFlatStatus() fiber.Handler
	GetFlatsByHouseID() fiber.Handler
	SubscribeToHouse() fiber.Handler
}

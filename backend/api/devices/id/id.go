package id

import (
	"github.com/gofiber/fiber/v2"
)

func SetupDeviceId(router fiber.Router) {
	router.Get("/devices/:id", getDeviceRoute())

	router.Delete("/devices/:id", deleteDeviceRoute())

	router.Get("/devices/:id/switch", switchRoute)
}

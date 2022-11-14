package id

import (
	"github.com/gofiber/fiber/v2"
)

func SetupDeviceId(router fiber.Router) {
	router.Get("/devices/:id", getDeviceRoute())
	router.Get("/devices/:id/state", getStateRoute())

	router.Delete("/devices/:id", deleteDeviceRoute())

	router.Post("/devices/:id/switch", switchRoute)
	router.Post("/devices/:id/light", lightRoute)
}

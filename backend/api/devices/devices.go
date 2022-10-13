package devices

import (
	"github.com/gofiber/fiber/v2"
)

func SetupDevices(router fiber.Router) {
	router.Get("/devices/list", listDevicesRoute())

	router.Get("/devices/:id", getDeviceRoute())

	router.Post("/devices/create", createDeviceRoute())

	router.Delete("/devices/:id", deleteDeviceRoute())
}

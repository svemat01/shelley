package devices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/api/devices/id"
)

func SetupDevices(router fiber.Router) {
	router.Get("/devices/list", listDevicesRoute())

	router.Post("/devices/create", createDeviceRoute())

	id.SetupDeviceId(router)
}

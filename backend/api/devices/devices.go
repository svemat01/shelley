package devices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
)

func SetupDevices(router fiber.Router, data *pkg.MainData) {
	router.Get("/devices/list", listDevicesRoute(data))

	router.Get("/devices/:id", getDeviceRoute(data))

	router.Post("/devices/create", createDeviceRoute(data))
}

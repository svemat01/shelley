package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/api/devices"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	devices.SetupDevices(v1)
}

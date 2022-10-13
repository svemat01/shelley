package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/api/devices"
	"github.com/svemat01/shelley/pkg"
)

func Setup(app *fiber.App, data *pkg.MainData) {
	v1 := app.Group("/api/v1")

	devices.SetupDevices(v1, data)
}

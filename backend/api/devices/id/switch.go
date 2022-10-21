package id

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
	"github.com/svemat01/shelley/shelly"
	"log"
)

type switchRouteData struct {
	On          *bool  `json:"on" validate:"required"`
	SwitchIndex string `json:"switch" validate:"required"`
}

func switchRoute(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "application/json" {
		return pkg.BadRequest("Content-Type must be application/json")
	}

	payload := new(switchRouteData)

	// Parse JSON body
	if err := c.BodyParser(payload); err != nil {
		return pkg.BadRequest("Invalid JSON")
	}

	log.Println(payload)

	// Validate payload
	if err := pkg.ValidateStruct(*payload); err != nil {
		return pkg.ValidatioError(err)
	}

	// Get device id from URL
	key := c.Params("id")

	exists, err := redisDB.ExistsDevice(key)

	if err != nil {
		return pkg.Unexpected("Failed")
	}

	if !exists {
		return pkg.NotFound("Device not found")
	}

	err = shelly.SetSwitchState(key, payload.SwitchIndex, *payload.On)

	if err != nil {
		return pkg.Unexpected("Failed to set switch state")
	}

	return c.SendStatus(fiber.StatusOK)
}

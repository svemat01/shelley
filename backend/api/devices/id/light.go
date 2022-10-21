package id

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
	"github.com/svemat01/shelley/shelly"
	"log"
)

type lightRouteData struct {
	On         *bool  `json:"on" validate:"required"`
	Brightness int    `json:"brightness"`
	LightIndex string `json:"light" validate:"required"`
}

func lightRoute(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType != "application/json" {
		return pkg.BadRequest("Content-Type must be application/json")
	}

	payload := new(lightRouteData)

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

	err = shelly.SetLightState(key, payload.LightIndex, *payload.On, payload.Brightness)

	if err != nil {
		return pkg.Unexpected("Failed to set light state")
	}

	return c.SendStatus(fiber.StatusOK)
}

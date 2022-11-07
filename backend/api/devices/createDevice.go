package devices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
	"strconv"
)

func createDeviceRoute() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// TODO check if device is valid type

		contentType := c.Get("Content-Type")

		if contentType != "application/json" {
			return pkg.BadRequest("Content-Type must be application/json")
		}

		payload := new(redisDB.Device)

		// Parse JSON body
		if err := c.BodyParser(payload); err != nil {
			return pkg.BadRequest("Invalid JSON")
		}

		// Validate payload
		if err := pkg.ValidateStruct(*payload); err != nil {
			return pkg.ValidatioError(err)
		}

		id, err := pkg.SonyFlake.NextID()

		if err != nil {
			return pkg.Unexpected("ID generation error")
		}

		err = redisDB.CreateDevice(strconv.FormatUint(id, 10), *payload)

		if err != nil {
			return pkg.Unexpected("Error creating device")
		}

		return c.JSON(fiber.Map{
			"device_id": id,
		})
	}
}

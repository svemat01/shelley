package devices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
)

type Device struct {
	Name string `json:"name" validate:"required" `
	Ip   string `json:"ip" validate:"required"`
}

func createDeviceRoute(data pkg.MainData) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")

		if contentType != "application/json" {
			return pkg.BadRequest("Content-Type must be application/json")
		}

		payload := new(Device)

		// Parse JSON body
		if err := c.BodyParser(payload); err != nil {
			return err
		}

		// Validate payload
		if err := pkg.ValidateStruct(*payload); err != nil {
			return pkg.ValidatioError(err)
		}

		return c.JSON(payload)
	}
}

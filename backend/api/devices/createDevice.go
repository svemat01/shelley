package devices

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
)

type Device struct {
	Name string `json:"name" validate:"required"`
	Ip   string `json:"ip"`
}

func createDeviceRoute(data pkg.MainData) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		contentType := c.Get("Content-Type")

		if contentType != "application/json" {
			return pkg.BadRequest("Content-Type must be application/json")
		}

		payload := new(Device)

		if err := c.BodyParser(payload); err != nil {
			return err
		}
		err := pkg.ValidateStruct[Device](*payload)
		fmt.Println(err)
		if err != nil {
			json, _ := json.Marshal(err)
			return pkg.BadRequest(string(json))
		}

		return c.JSON(payload)
	}
}

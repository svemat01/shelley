package id

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
)

func getStateRoute() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		key := c.Params("id")

		exists, err := redisDB.ExistsDevice(key)

		if err != nil {
			return pkg.Unexpected("Failed")
		}

		if !exists {
			return pkg.NotFound("Device not found")
		}

		val, err := redisDB.GetDeviceState(key)

		switch {
		case err == redis.Nil:
			return pkg.NotFound("Device doesn't exist")

		case err != nil:
			return pkg.Unexpected("Failed")
		}

		return c.JSON(val)
	}
}

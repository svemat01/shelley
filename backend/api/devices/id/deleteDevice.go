package id

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
)

func deleteDeviceRoute() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		key := c.Params("id")
		exists, err := redisDB.ExistsDevice(key)

		switch {
		case err == redis.Nil || !exists:
			return pkg.NotFound("Device not found")

		case err != nil:
			return pkg.Unexpected("Failed")
		}

		err = redisDB.DeleteDevice(key)

		if err != nil {
			return pkg.Unexpected("Failed")
		}

		return c.SendStatus(200)
	}
}

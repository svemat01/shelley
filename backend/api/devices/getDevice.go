package devices

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
)

func getDeviceRoute(data pkg.MainData) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		key := c.Params("id")
		val, err := redisDB.GetDevice(data.Redis, data.RedisContext, key)

		switch {
		case err == redis.Nil:
			return pkg.NotFound("Device doesn't exist")

		case err != nil:
			return pkg.Unexpected("Failed")
		}

		return c.JSON(val)
	}
}

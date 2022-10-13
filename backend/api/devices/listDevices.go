package devices

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
)

func listDevicesRoute(data *pkg.MainData) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		val, err := redisDB.GetDevices(data.Redis, data.RedisContext)
		if err != nil {
			if err == redis.Nil {
				return pkg.NotFound("[] nil")
			}
			fmt.Println(err)
			return pkg.Unexpected("Redis connection failed.")
		}

		switch {
		case err == redis.Nil:
			return pkg.NotFound("[] nil")

		case err != nil:
			return pkg.Unexpected("Failed")

		case val == nil:
			return c.SendString("[]")
		}

		return c.JSON(val)
	}
}

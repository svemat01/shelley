package api

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
)

func Setup(app *fiber.App, data pkg.MainData) {
	v1 := app.Group("/api/v1")

	v1.Get("/devices/list", func(c *fiber.Ctx) error {
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
	})

	v1.Get("/devices/:id", func(c *fiber.Ctx) error {
		key := c.Params("id")
		val, err := redisDB.GetDevice(data.Redis, data.RedisContext, key)

		switch {
		case err == redis.Nil:
			return pkg.NotFound("Device doesn't exist")

		case err != nil:
			return pkg.Unexpected("Failed")
		}

		return c.JSON(val)
	})
}

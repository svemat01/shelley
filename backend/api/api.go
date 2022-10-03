package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var count = 0

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		count++
		return c.SendString(fmt.Sprintf("Hello, World! %d", count))
	})
}

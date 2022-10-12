package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/svemat01/shelley/pkg"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(cors.New())
	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New())
	}
}

func Create() *fiber.App {

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*pkg.APIError); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*pkg.APIErrorData); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(pkg.APIError{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(pkg.APIError{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
	})

	setupMiddlewares(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}

func Listen(app *fiber.App) error {

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(pkg.APIError{Status: 404, Code: "not-found", Message: "Not Found"})
	})

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	return app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))
}

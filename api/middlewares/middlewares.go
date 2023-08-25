package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MakeMiddlewares(app *fiber.App) {
	// app.Use(logger.New())
	app.Use(recover.New())
	app.Use(setContentType)
}

func setContentType(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "application/json")
	return c.Next()
}

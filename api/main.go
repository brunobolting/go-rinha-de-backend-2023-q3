package main

import (
	"fmt"

	"github.com/brunobolting/go-rinha-backend/config/env"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%s", env.PORT))
}

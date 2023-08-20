package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func home(c *fiber.Ctx) error {
	return c.JSON("Hey World")
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", home)

	app.Listen(":3000")
}

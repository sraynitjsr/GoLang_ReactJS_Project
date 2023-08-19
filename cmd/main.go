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

	app.Use(cors.New(cors.Config{
		Next:             nil,
		AllowOriginsFunc: nil,
		AllowOrigins:     "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	app.Get("/", home)
	app.Listen(":3000")
}

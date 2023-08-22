package main

import "github.com/gofiber/fiber/v2"

func home(c *fiber.Ctx) error{
	return c.SendString("Hello World")
}

func main() {
  app := fiber.New()

  app.Get("/", home)

  app.Listen(":3000")
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ResponseData struct {
	Message string `json:"message"`
}

func helloHandler(c *fiber.Ctx) error {
	responseData := ResponseData{
		Message: "Hello, World!",
	}

	return c.JSON(responseData)
}

func main() {
	app := fiber.New()

	// Configure CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://gofiber.io, https://gofiber.net",
		AllowMethods:     "GET",
		AllowHeaders:     "Origin, Content-Type",
		AllowCredentials: true,
	}))

	// Define the GET endpoint
	app.Get("/", helloHandler)

	// Start the Fiber app
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

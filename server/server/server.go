package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sraynitjsr/server/handler"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", handler.Home)
	app.Get("/getAllData", handler.GetAllData)
	app.Get("/getNEByType", handler.GetNEByType)
	app.Get("/getByBuildVersion", handler.GetByBuildVersion)
	app.Get("/getByNETypeAndBuildVersion", handler.GetByNETypeAndBuildVersion)
	app.Get("/getBySupportedConfig", handler.GetBySupportedConfig)

	app.Listen(":3000")
}

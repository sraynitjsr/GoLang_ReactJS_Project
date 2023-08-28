package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sray/cmd/config"
	"github.com/sray/cmd/handler"
)

func main() {
	app := fiber.New()

	app.Get("/", handler.HelloHandler)

	app.Get("/getAllStudents", handler.GetAllStudents)

	app.Get("/getStudentByRollNumber", handler.GetStudentByRollNumber)

	app.Get("/getStudentByName/:name?", handler.GetStudentByName)

	app.Listen(config.AppPort)
}

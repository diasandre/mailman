package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: format,
	}))

	app.Use(recover.New())

	app.Post("/", handler)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

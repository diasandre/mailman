package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Post("/", handler)

	app.Use(logger.New())

	app.Listen(":3000")
}

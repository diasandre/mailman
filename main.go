package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func main() {
	app := fiber.New()

	InitLogger()

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

func InitLogger() { zapLogger, _ = zap.NewProduction() }

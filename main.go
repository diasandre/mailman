package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

var projectID = "YOUR_PROJECT_ID"
var topicId = "values"

var restClient = resty.New()
var client *pubsub.Client
var zapLogger *zap.Logger

func main() {
	app := fiber.New()

	app.Use(recover.New())

	initLogger(app)
	initPubSub()

	setupRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func initLogger(app *fiber.App) {
	zapLogger, _ = zap.NewProduction()

	app.Use(logger.New(logger.Config{
		Format: format,
	}))
}

func initPubSub() {
	ctx := context.Background()

	var err error
	client, err = pubsub.NewClient(ctx, projectID)

	if err != nil {
		logError(err)
	}

	sub := client.Subscription(topicId)
	err = sub.Receive(ctx, subscribe)

	if err != nil {
		logError(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It's UP")
	})
	app.Post("/", publish)
}

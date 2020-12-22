package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"github.com/gofiber/fiber/v2"
)

func publish(c *fiber.Ctx) error {
	topic := client.Topic(topicId)
	res := topic.Publish(context.Background(), &pubsub.Message{
		Data: c.Body(),
	})

	_, err := res.Get(context.Background())

	if err != nil {
		logError(err)
		return consumeNow(c)
	}

	return c.SendStatus(fiber.StatusOK)
}

func consumeNow(c *fiber.Ctx) error {
	event, err := parser(c)
	if err != nil {
		logError(err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = consume(event)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func subscribe(_ context.Context, m *pubsub.Message) {
	event, err := fromJson(m.Data)
	if err != nil {
		logError(err)
	} else {
		err = consume(event)
		if err == nil {
			m.Ack()
		}
	}
}

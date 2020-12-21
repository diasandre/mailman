package main

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func consume(event Event, c *fiber.Ctx) error {
	logReceivedEvent(event)
	switch eventType := event.Type; eventType {
	default:
		return defaultConsumer(event, c)
	}
}

func logReceivedEvent(event Event) {
	zapLogger.Info(
		receivedEventLog,
		zap.String("eventType", string(event.Type)),
		zap.String("payload", event.Payload))
}

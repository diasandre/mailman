package main

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func consume(event Event, c *fiber.Ctx) error {
	logReceivedEvent(event)
	switch eventType := event.Type; eventType {
	case example:
		return exampleConsumer(event, c)
	default:
		return event.handleUnsupportedEventType(c)
	}
}

func (event Event) handleUnsupportedEventType(c *fiber.Ctx) error {
	err := errors.New(fmt.Sprintf(unsupportedEventType, event.Type))
	zapLogger.Error(err.Error())
	return c.Status(fiber.StatusBadRequest).SendString(err.Error())
}

func logReceivedEvent(event Event) {
	zapLogger.Info(
		receivedEventLog,
		zap.String("eventType", string(event.Type)),
		zap.String("payload", event.Payload))
}

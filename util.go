package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

var format = "${pid} - ${status} - ${method} - ${path} - ${body}\n"

var responseLog = "response"
var receivedEventLog = "received event"
var unsupportedEventType = "unsupported event: %s"

func logError(err error) {
	zapLogger.Error(err.Error())
}

func parser(c *fiber.Ctx) (Event, error) {
	var event Event
	err := c.BodyParser(&event)
	return event, err
}

func fromJson(data []byte) (Event, error) {
	var event Event
	err := json.Unmarshal(data, &event)
	return event, err
}

package main

import (
	"github.com/gofiber/fiber/v2"
)

var format = "${pid} - ${status} - ${method} - ${path} - ${body}\n"

var successResponseLog = "success response"
var receivedEventLog = "received event"

func logError(err error) {
	zapLogger.Error(err.Error())
}

func parser(c *fiber.Ctx) (Event, error) {
	var event Event
	err := c.BodyParser(&event)
	return event, err
}

package main

import (
	"github.com/gofiber/fiber"
)

var format = "${pid} - ${status} - ${method} - ${path} - ${body}\n"

func parser(c *fiber.Ctx) (Event, error) {
	var event Event
	err := c.BodyParser(&event)
	return event, err
}

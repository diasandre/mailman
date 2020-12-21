package main

import (
	"github.com/gofiber/fiber/v2"
)

func handler(c *fiber.Ctx) error {
	event, err := parser(c)
	if err != nil {
		logError(err)
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return consume(event, c)
}

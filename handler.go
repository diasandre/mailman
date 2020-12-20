package main

import (
	"github.com/gofiber/fiber"
	"log"
)

func handler(c *fiber.Ctx) error {
	event, err := parser(c)
	if err != nil {
		log.Printf(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	consume(event)

	return c.SendString(event.Payload)
}

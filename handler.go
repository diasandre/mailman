package main

import (
	"encoding/json"
	"github.com/gofiber/fiber"
)

type Event struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

func fromJson(data []byte) (Event, error) {
	var event Event
	err := json.Unmarshal(data, &event)
	return event, err
}

func handler(c *fiber.Ctx) error {
	event, err := fromJson(c.Body())

	if err != nil {
		c.SendStatus(500)
		return c.SendString(err.Error())
	}

	return c.SendString(event.Payload)
}

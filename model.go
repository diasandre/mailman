package main

type EventType string

type Event struct {
	Type    EventType `json:"type"`
	Payload string    `json:"payload"`
}

const (
	example EventType = "EXAMPLE"
)

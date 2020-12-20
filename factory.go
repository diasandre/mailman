package main

import (
	"fmt"
)

func consume(event Event) {
	switch eventType := event.Type; eventType {
	case UpdateCase:
		fmt.Printf("creating consumer to handle %s.\n", event.Type)
	default:
		fmt.Printf("%s not supported.\n", event.Type)
	}
}

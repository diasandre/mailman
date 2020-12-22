package main

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
)

func consume(event Event) error {
	logReceivedEvent(event)
	switch eventType := event.Type; eventType {
	case example:
		return exampleConsumer(event)
	default:
		return event.handleUnsupportedEventType()
	}
}

func (event Event) handleUnsupportedEventType() error {
	err := errors.New(fmt.Sprintf(unsupportedEventType, event.Type))
	zapLogger.Error(err.Error())
	return err
}

func logReceivedEvent(event Event) {
	zapLogger.Info(
		receivedEventLog,
		zap.String("eventType", string(event.Type)),
		zap.String("payload", event.Payload))
}

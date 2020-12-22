package main

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

func exampleConsumer(event Event) error {
	resp, err := request(event.Payload, "https://example:8080/api")
	return responseHandler(resp, event, err)
}

func request(payload string, url string) (*resty.Response, error) {
	return restClient.R().SetBody(payload).Post(url)
}

func responseHandler(resp *resty.Response, event Event, err error) error {
	logResponse(event, resp)
	if err != nil {
		logError(err)
	}
	return err
}

func logResponse(event Event, resp *resty.Response) {
	zapLogger.Info(
		responseLog,
		zap.String("eventType", string(event.Type)),
		zap.String("payload", event.Payload),
		zap.Int("statusCode", resp.StatusCode()),
		zap.ByteString("body", resp.Body()))
}

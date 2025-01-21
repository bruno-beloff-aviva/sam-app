package main

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var logger = slog.Default()
var level = slog.LevelInfo

func New200Response(message string) events.APIGatewayProxyResponse {
	payload := map[string]string{"message": "Canary deployments 3 - " + message}
	body, _ := json.Marshal(payload)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logger.Info("got request")

	var message string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		message = "Hello, world!!"
	} else {
		message = fmt.Sprintf("Hello client %s!!", sourceIP) // request.RequestContext.APIID,
	}

	return New200Response(message), nil
}

func main() {
	slog.SetLogLoggerLevel(level)

	lambda.Start(handler)
}

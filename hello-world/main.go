package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func New200Response(message string) events.APIGatewayProxyResponse {
	payload := map[string]string{"message": "Canary deployments: " + message}
	body, _ := json.Marshal(payload)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var message string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		message = "Hello, world!!"
	} else {
		message = fmt.Sprintf("Hello api: %s client: %s!!", request.RequestContext.APIID, sourceIP)
	}

	return New200Response(message), nil
}

func main() {
	lambda.Start(handler)
}

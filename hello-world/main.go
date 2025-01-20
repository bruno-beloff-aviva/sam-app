package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func New200Response(body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: 200,
	}
}

//       body: JSON.stringify({
//	message: "hello my friend",
//}),

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	if sourceIP == "" {
		greeting = "Hello, world!!\n"
	} else {
		greeting = fmt.Sprintf("Hello api: %s client: %s!!\n", request.RequestContext.APIID, sourceIP)
	}

	return New200Response(greeting), nil
}

func main() {
	lambda.Start(handler)
}

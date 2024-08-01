package main

import (
	"fmt"
	"os"
	"sam-json-logs/logs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	logs.Request = request

	hello := "Hello from json-logs"

	logs.Info("ENV", os.Environ())

	logs.Warn("WRONG_DATA", map[string]any{"id": "some-uuid"})

	logs.Debug("SOME_INFO", map[string]any{"id": "some-uuid", "jere": "there"})

	fmt.Println(hello)
	return events.APIGatewayProxyResponse{Body: hello, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

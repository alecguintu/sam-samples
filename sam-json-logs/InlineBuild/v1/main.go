package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	hello := "Hello from InlineBuild"

	fmt.Println(hello)
	return events.APIGatewayProxyResponse{Body: hello, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

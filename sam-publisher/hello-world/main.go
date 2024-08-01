package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	hello := "Hello from publisher"

	// _, err := SNSPublish(hello, os.Getenv("PUBLISH_TOPIC"))
	// if err != nil {
	// 	fmt.Printf("SNS_ERR: %+v", err)
	// 	return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	// }

	fmt.Println(hello)
	return events.APIGatewayProxyResponse{Body: hello, StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}

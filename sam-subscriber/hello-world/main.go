package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.SQSEvent) (events.SQSEventResponse, error) {
	hello := "Hello from subscriber"
	var failedItems []events.SQSBatchItemFailure

	for _, record := range request.Records {
		fmt.Printf("RECORD: %+v\n", record)

		// if err {
		// 	failedItems = append(failedItems, events.SQSBatchItemFailure{
		// 		ItemIdentifier: record.MessageId,
		// 	})
		// 	continue
		// }
	}

	fmt.Println(hello)
	return events.SQSEventResponse{
		BatchItemFailures: failedItems,
	}, nil
}

func main() {
	lambda.Start(handler)
}

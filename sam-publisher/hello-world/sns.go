package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/scrambledeggs/booky-go-common/logger"
)

func SNSPublish(message string, topic string) (*sns.PublishOutput, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	snsClient := sns.NewFromConfig(cfg)

	t := []string{"arn", "aws", "sns", os.Getenv("AMZ_REGION"), os.Getenv("AMZ_ACCOUNT_ID"), topic}
	topicArn := strings.Join(t, ":")
	fmt.Printf("TOPIC: %+v", topicArn)

	out, err := snsClient.Publish(context.TODO(),
		&sns.PublishInput{
			TopicArn: aws.String(topicArn),
			Message:  aws.String(message),
		})
	if err != nil {
		return nil, err
	}

	logger.Infof("[SNS_PUBLISH] Published to SNS successfully. messageId: %v", *out.MessageId)
	return out, nil
}

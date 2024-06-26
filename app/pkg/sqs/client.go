package sqsclient

import (
	"bytes"
	"context"
	"encoding/json"
	database "learn-sqs/app/pkg/database/model"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/samber/lo"
)

type Sqs struct {
	client   *sqs.Client
	queueURL string
}

func Init() (*Sqs, error) {
	cfg, err := initAwsConfig()
	if err != nil {
		return nil, err
	}

	client := sqs.NewFromConfig(*cfg)

	return &Sqs{
		client:   client,
		queueURL: os.Getenv("QUEUE_URL"),
	}, nil
}

func initAwsConfig() (*aws.Config, error) {
	// TODO: Deprecatedな関数の置き換え
	resolver := aws.EndpointResolverWithOptionsFunc(
		func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
			if os.Getenv("ENV") != "production" {
				//nolint:exhaustruct
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           os.Getenv("QUEUE_URL"),
					SigningRegion: os.Getenv("AWS_REGION"),
				}, nil
			}
			//nolint:exhaustruct
			return aws.Endpoint{}, &aws.EndpointNotFoundError{}
		})

	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithEndpointResolverWithOptions(resolver))
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (s Sqs) SendMessage(ctx context.Context, message database.Message) error {
	messageBody := bytes.NewBuffer(
		lo.Must(json.Marshal(message)),
	).String()

	params := &sqs.SendMessageInput{
		MessageBody: &messageBody,
		QueueUrl:    &s.queueURL,
	}

	_, err := s.client.SendMessage(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

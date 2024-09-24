package main

import (
	"context"
	"fmt"

	"github.com/Jardielson-s/lambdas-go/src/infra/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, s3Event events.S3Event) (any, error) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, key = %s \n", record.EventSource, record.EventTime, s3.Bucket, s3.Object.Key)
	}
	return services.Process_csv_service()
}

func main() {
	lambda.Start((handler))
}

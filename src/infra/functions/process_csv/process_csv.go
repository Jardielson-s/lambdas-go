package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config/types"
	"github.com/Jardielson-s/lambdas-go/src/infra/services"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, s3Event events.S3Event) (any, error) {
	log.Println("handler - processing: ", s3Event)
	log.Println("Event:", s3Event.Records)
	input := make(map[string]string)
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, key = %s \n", record.EventSource, record.EventTime, s3.Bucket, s3.Object.Key)
		input["Bucket"] = s3.Bucket.Name
		input["Key"] = s3.Object.Key
	}
	services.ProcessCsvService(types.GetFileInput{
		Bucket: input["Bucket"],
		Key:    input["Key"],
	})
	log.Println("handler - processed")
	return "handler - processed", nil
}

func main() {
	lambda.Start((handler))
}

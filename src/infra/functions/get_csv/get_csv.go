package main

import (
	"context"

	"github.com/Jardielson-s/lambdas-go/src/infra/services"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context) (any, error) {
	println(ctx)
	return services.Process_csv_service()
}

func main() {
	lambda.Start((handler))
}

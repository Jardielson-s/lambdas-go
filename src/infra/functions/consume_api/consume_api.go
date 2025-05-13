package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Jardielson-s/lambdas-go/src/domain/structs/users"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SQSMessage struct {
	Data []users.User `json:"data"`
}

func handler(_ context.Context, sqsvent events.SQSEvent) (any, error) {
	url := os.Getenv("TRANSFER_X_API")

	log.Println("handler - processing: ", sqsvent)
	log.Println("Event:", sqsvent.Records)
	for _, record := range sqsvent.Records {
		body := record.Body
		fmt.Printf(`Body: %s`, body)
		var sqsMessage SQSMessage
		err := json.Unmarshal([]byte(record.Body), &sqsMessage)
		if err != nil {
			log.Printf("Failed to unmarshal SQS message body: %v", err)
			return nil, err
		}
		requestBody, err := json.Marshal(sqsMessage.Data)
		if err != nil {
			log.Printf("Failed to marshal request body: %v", err)
			return nil, err
		}

		send_api_request(url, requestBody, "POST")

	}
	log.Println("handler - processed")
	return "handler - processed", nil
}

func main() {
	url := os.Getenv("TRANSFER_X_API")
	fmt.Println((url))
	lambda.Start((handler))
}

func send_api_request(url string, data []byte, verb string) {
	req, err := http.NewRequest(verb, url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	req.Header.Add("token", "key")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println((res))
	defer res.Body.Close()
	body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println(res.Status, "Response Body: ", string(body))
}

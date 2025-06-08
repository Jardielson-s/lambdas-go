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

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SQSMessage struct {
	Entity    string `json:"entity"`
	Data      string `json:"data"`
	Service   string `json:"service"`
	Operation string `json:"operation"`
}

type SendRequest struct {
	Data []string `json:"data"`
}

type APIRequest struct {
	Data []map[string]interface{} `json:"data"`
}

func handler(_ context.Context, sqsvent events.SQSEvent) (any, error) {
	log.Println("handler - processing: ", sqsvent)
	for _, record := range sqsvent.Records {
		body := record.Body
		fmt.Printf(`Body: %s`, body)
		var sqsMessage SQSMessage
		var data []map[string]interface{}
		err := json.Unmarshal([]byte(record.Body), &sqsMessage)

		if err != nil {
			log.Printf("Failed to unmarshal SQS message body: %v", err)
			return nil, err
		}
		json.Unmarshal([]byte(sqsMessage.Data), &data)

		request := APIRequest{
			Data: data,
		}
		fmt.Println("Service: ", sqsMessage.Service)
		fmt.Println("Enitity: ", data)
		var url string
		var verb string
		if sqsMessage.Service == "user-ms" {
			url = os.Getenv("TRANSFER_X_API") + sqsMessage.Entity + "/upsert"
			if sqsMessage.Operation == "I" {
				verb = "POST"
			} else {
				url = os.Getenv("TRANSFER_X_API") + sqsMessage.Entity
				verb = "PATCH"
			}
		} else {
			url = os.Getenv("USER_MS_API") + sqsMessage.Entity
			verb = "PATCH"
		}

		fmt.Println(url)

		payload, _ := json.Marshal(request)
		send_api_request(url, payload, verb)

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
	fmt.Println("Response Body: ", string(body))
}

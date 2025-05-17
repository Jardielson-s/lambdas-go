package services

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config"
	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config/types"
	sqsconfig "github.com/Jardielson-s/lambdas-go/src/infra/sqs_config"
)

type MessageBody struct {
	Entity string `json:"entity"`
	Data   string `json:"data"`
}

func ProcessCsvService(input types.GetFileInput) (response string, err any) {
	log.Println("Process_csv_service - processing: ", input)
	pathEntity := strings.Split(input.Key, "/")

	var entity string = pathEntity[2]
	log.Println("Entity: ", entity)
	rows, err := s3_config.GetRows(input, s3_config.GetFile, s3_config.GetHeaders, entity)
	if err != nil {
		log.Println("Process_csv_service - processed error:", err)
		return "", err
	}
	// log.Println("Process_csv_service - processed:", string(rows))
	log.Println("Sending rows to sqs")
	queueUrl, err := sqsconfig.GetQueue()
	if err != nil {
		log.Println("Sended Queue Error: ", err)
		return "", err
	}
	message := MessageBody{
		Entity: entity,
		Data:   string(rows),
	}
	body, err := json.Marshal(message)
	sqsconfig.SendMessage(queueUrl, string(body))
	log.Println("Sended rows to Sqs")
	return "Sended rows to Sqs", nil
}

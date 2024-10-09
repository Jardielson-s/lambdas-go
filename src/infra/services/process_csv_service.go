package services

import (
	"log"

	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config"
	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config/types"
	sqsconfig "github.com/Jardielson-s/lambdas-go/src/infra/sqs_config"
)

func ProcessCsvService(input types.GetFileInput) {
	log.Println("Process_csv_service - processing: ", input)
	rows, err := s3_config.GetRows(input, s3_config.GetFile, s3_config.GetHeaders)
	if err != nil {
		log.Println("Process_csv_service - processed error:", err)
	}
	log.Println("Process_csv_service - processed:", string(rows))
	log.Println("Sending rows to sqs")
	queueUrl, err := sqsconfig.GetQueue()
	if err != nil {
		log.Println("Sended Queue Error: ", err)
		return
	}
	sqsconfig.SendMessage(queueUrl, string(rows))
	log.Println("Sended rows to Sqs")
}

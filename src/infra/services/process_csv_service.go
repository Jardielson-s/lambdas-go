package services

import (
	"log"

	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config"
	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config/types"
)

func ProcessCsvService(input types.GetFileInput) {
	log.Println("Process_csv_service - processing: ", input)
	rows, err := s3_config.GetRows(input, s3_config.GetFile, s3_config.GetHeaders)
	if err != nil {
		log.Println("Process_csv_service - processed error:", err)
	} else {
		log.Println("Process_csv_service - processed:", string(rows))
	}
}

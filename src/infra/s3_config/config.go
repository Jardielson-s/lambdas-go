package s3_config

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"

	"github.com/Jardielson-s/lambdas-go/src/domain/structs/users"
	"github.com/Jardielson-s/lambdas-go/src/infra/s3_config/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Session *s3.S3
)

func init() {
	log.Println("connecting with s3")
	region := os.Getenv("AWS_REGION")
	s3Session = s3.New(session.Must(session.NewSession((&aws.Config{
		Region: aws.String(region),
	}))))
	log.Println("connected with s3")
}

func GetFile(input types.GetFileInput) *s3.GetObjectOutput {
	rawObject, err := s3Session.GetObject(&s3.GetObjectInput{Bucket: aws.String(input.Bucket), Key: aws.String(input.Key)})
	if err != nil {
		log.Println("Error to get file")
		log.Println(err)
	}
	log.Println("File data: ", rawObject)

	return rawObject
}

func GetHeaders(records [][]string, entity string) (response []byte, operation string, err any) {
	var data []map[string]any
	var columnNames []string

	switch entity {
	case "users":
		columnNames = users.UserColumnsValues
	}
	var opr string
	for _, record := range records {
		row := make(map[string]any)
		if len(record) > 0 && (record[0] == "U" || record[0] == "I" || record[0] == "D") {
			opr = record[0]
			record = record[1:]
		}
		for i, value := range record {
			row[string(columnNames[i])] = value
		}
		if opr == "I" || opr == "U" {
			if !(row["integration_id"] != nil && opr == "I") {
				data = append(data, row)
			}
		}
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, "", err
	}
	return jsonData, opr, nil
}

func GetRows(input types.GetFileInput, getFile types.GetFile, getHeaders types.GetHeaders, entity string) (response []byte, operation string, err any) {
	rowObject := getFile(input)
	data, err := csv.NewReader(rowObject.Body).ReadAll()
	if err != nil {
		return nil, "", err
	}
	log.Println("Processing data:", data)
	rows, operation, err := getHeaders(data, entity)
	return rows, operation, err
}

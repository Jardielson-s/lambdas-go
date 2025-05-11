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

func GetHeaders(records [][]string, entity string) (response []byte, err any) {
	var data []map[string]any
	var columnNames []string

	switch entity {
	case "users":
		columnNames = users.UserColumnsValues
	}

	for _, record := range records {
		row := make(map[string]any)
		for i, value := range record {
			row[string(columnNames[i])] = value
		}
		data = append(data, row)
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func GetRows(input types.GetFileInput, getFile types.GetFile, getHeaders types.GetHeaders, entity string) (response []byte, err any) {
	rowObject := getFile(input)
	data, err := csv.NewReader(rowObject.Body).ReadAll()
	if err != nil {
		return nil, err
	}
	log.Println("Processing data:", data)
	rows, err := getHeaders(data, entity)
	return rows, err
}

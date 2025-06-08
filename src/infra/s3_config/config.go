package s3_config

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
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

func transformUserRow(row map[string]any) map[string]any {
	transformed := make(map[string]any)
	if row["integration_id"] != "" {
		transformed["_id"] = row["integration_id"]
	}
	if row["external_application"] != "" {
		transformed["externalApplication"] = row["external_application"]
	}
	transformed["externalId"] = row["id"]
	transformed["name"] = row["name"]
	transformed["ein"] = row["ein"]
	transformed["age"] = 18
	transformed["email"] = row["email"]
	transformed["password"] = row["password"]
	address := map[string]any{
		"street":        row["address"],
		"addressNumber": row["addressNumber"],
	}
	address["postalCode"] = row["postalCode"]
	transformed["address"] = address
	return transformed
}

func GetHeaders(records [][]string, entity string) (response []byte, operation string, err any) {
	var columnNames []string

	switch entity {
	case "users":
		columnNames = users.UserColumnsValues
	default:
		return nil, "", fmt.Errorf("entidade '%s' não suportada", entity)
	}

	var data []map[string]any
	var opr string

	for _, record := range records {
		if len(record) == 0 {
			continue
		}

		row := make(map[string]any)

		if record[0] == "U" || record[0] == "I" || record[0] == "D" {
			opr = record[0]
			record = record[1:]
		}

		for i, value := range record {
			if i < len(columnNames) {
				row[columnNames[i]] = value
			}
		}
		fmt.Println("Row ", row)
		if entity == "users" && (opr != "D") {
			if !((row["external_application"] == "true" || row["external_application"] == true) && opr == "I") {
				transformed := transformUserRow(row)
				data = append(data, transformed)
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

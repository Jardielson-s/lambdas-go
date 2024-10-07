package sqsconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	sqsSession *sqs.SQS
)

func init() {
	log.Println("connecting with sqs")
	region := os.Getenv("AWS_REGION")
	sqsSession = sqs.New(session.Must(session.NewSession((&aws.Config{
		Region: aws.String(region),
	}))))
	log.Println("connected with sqs")
}

func GetQueue() (queueUrl string, err any) {
	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String("sync-" + os.Getenv("ENV")),
	}
	response, err := sqsSession.GetQueueUrl(input)
	if err != nil {
		return "", err
	}
	return string(*response.QueueUrl), nil
}

func SendMessage(queueUrl string, message string) (*sqs.SendMessageOutput, error) {
	sendMessageInput := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(message),
	}
	result, err := sqsSession.SendMessage(sendMessageInput)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, nil
}

// func main() {
// 	queueUrl, _ := GetQueue()
// 	SendMessage(queueUrl, "hello queue")
// }

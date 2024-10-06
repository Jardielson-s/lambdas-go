package types

import "github.com/aws/aws-sdk-go/service/s3"

type GetFileInput struct {
	Bucket string
	Key    string
}
type GetFile func(input GetFileInput) *s3.GetObjectOutput
type GetHeaders func(records [][]string) (response []byte, err any)

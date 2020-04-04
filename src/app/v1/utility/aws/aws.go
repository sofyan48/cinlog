package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AWS ...
type AWS struct{}

// AWSHandler ...
func AWSHandler() *AWS {
	return &AWS{}
}

// AWSInterface ...
type AWSInterface interface{}

func initialize() *aws.Config {
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_ACCESS_SECRET"), "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(os.Getenv("AWS_ACCESS_AREA")).WithCredentials(creds)
	return cfgAws
}

// S3Session ...
func (svc *AWS) S3Session() *s3.S3 {
	cfg := initialize()
	return s3.New(session.New(), cfg)
}

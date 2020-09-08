package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/masgari/aws-custom-resources/internal"
)

var registery *internal.CustomResourceRegistery

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	registery = internal.New(sess)
}

func main() {
	lambda.Start(cfn.LambdaWrap(registery.HandleCustomResources))
}

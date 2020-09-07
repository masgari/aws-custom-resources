package main

import (
	//"context"
	//"fmt"
	//"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/masgari/aws-custom-resources/internal"
)

// func echoResource(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
// 	body, _ := json.Marshal(event)

// 	fmt.Printf("%s\n", string(body))

// 	v, _ := event.ResourceProperties["Echo"].(string)

// 	data = map[string]interface{}{
// 		"Echo": v,
// 	}

// 	return
// }


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

package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type apigwCustomResources struct {
	apigw     *apigateway.APIGateway
	ssmClient *ssm.SSM
}

func newApiGatewayResources(sess *session.Session) *apigwCustomResources {
	return &apigwCustomResources{
		apigw:     apigateway.New(sess),
		ssmClient: ssm.New(sess),
	}
}

func (a *apigwCustomResources) handleApiKey(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	switch event.RequestType {
	case "Create":
		return a.createApiKeyValueResource(ctx, event)
	case "Delete":
		return a.deleteApiKeyValueResource(ctx, event)
	default:
		err = fmt.Errorf("Unhandled request type: %s for %s", event.RequestType, ApiKeyValue)
	}
	return
}

func (a *apigwCustomResources) createApiKeyValueResource(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	keyId, err := strProperty(event, "ApiKeyId")
	if err != nil {
		return
	}

	ssmParamPath, err := strProperty(event, "SsmParamPath")
	if err != nil {
		return
	}
	physicalResourceID = ssmParamPath

	key, err := a.apigw.GetApiKey(&apigateway.GetApiKeyInput{
		ApiKey:       aws.String(keyId),
		IncludeValue: aws.Bool(true),
	})
	if err != nil {
		return
	}
	_, err = a.ssmClient.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(ssmParamPath),
		Value:     key.Value,
		Type:      aws.String(ssm.ParameterTypeSecureString),
		Overwrite: aws.Bool(boolPropertyWithDefault(event, "Overwrite", true)),
	})
	return
}

// deletes the ssm parameter created for this api-key value
func (a *apigwCustomResources) deleteApiKeyValueResource(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	ssmParamPath, err := strProperty(event, "SsmParamPath")
	if err != nil {
		return
	}
	physicalResourceID = ssmParamPath
	_, err = a.ssmClient.DeleteParameter(&ssm.DeleteParameterInput{
		Name: aws.String(ssmParamPath),
	})
	return
}

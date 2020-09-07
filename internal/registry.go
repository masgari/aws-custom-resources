package internal

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-lambda-go/cfn"
)


const (
	// resouce types
	ApiKeyValue = "Custom::ApiKeyValue"
)

type CustomResourceRegistery struct {
	apigw *apigwCustomResources
}

func New(sess *session.Session) *CustomResourceRegistery {
	return & CustomResourceRegistery{
		apigw: newApiGatewayResources(sess),
	}
}

func (h *CustomResourceRegistery) HandleCustomResources(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	switch event.ResourceType {
	case ApiKeyValue:
		return h.apigw.handleApiKey(ctx, event)
	default:
		err = fmt.Errorf("Unregistered resource type: %s", event.ResourceType)
	}
	return
}

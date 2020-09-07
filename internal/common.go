package internal

import (
	"fmt"
	"github.com/aws/aws-lambda-go/cfn"

)

func strProperty(event cfn.Event, propertyName string) (string, error) {
	if val, ok := event.ResourceProperties[propertyName]; ok {
		return val.(string), nil
	}
	return "", fmt.Errorf("Missing property %s", propertyName)
}

func boolPropertyWithDefault(event cfn.Event, propertyName string, defaultVal bool) bool {
	if val, ok := event.ResourceProperties[propertyName]; ok {
		return val.(bool)
	}
	return defaultVal
}


package internal

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/aws/aws-lambda-go/cfn"
)

func TestStrProperty(t *testing.T) {
	props := make(map[string]interface{})
	props["StrProperty"] = "value-1"
	event := cfn.Event {
		ResourceProperties: props,
	}

	val, err := strProperty(event, "StrProperty")
	assert.Equal(t, "value-1", val)
	assert.Nil(t, err)

	val, err = strProperty(event, "MissingStrProperty")
	assert.Equal(t, "", val)
	assert.NotNil(t, err)
}

func TestBoolPropertyWithDefault(t *testing.T) {
	props := make(map[string]interface{})
	props["BoolProperty"] = true
	event := cfn.Event {
		ResourceProperties: props,
	}

	val := boolPropertyWithDefault(event, "BoolProperty", false)
	assert.Equal(t, true, val)

	val = boolPropertyWithDefault(event, "MissingBoolProperty", true)
	assert.Equal(t, true, val)
}
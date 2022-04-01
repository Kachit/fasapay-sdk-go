package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Resource_NewResourceAbstract(t *testing.T) {
	config := BuildStubConfig()
	transport := NewHttpTransport(config, nil)
	resource := NewResourceAbstract(transport, config)
	assert.NotEmpty(t, resource)
}

func Test_Resource_ResourceAbstract_BuildAuthRequestParams(t *testing.T) {
	dt := BuildStubDateTime()
	config := BuildStubConfig()
	transport := NewHttpTransport(config, nil)
	resource := NewResourceAbstract(transport, config)
	result := resource.buildAuthRequestParams(dt)
	assert.NotEmpty(t, result)
	assert.Equal(t, config.ApiKey, result.ApiKey)
	assert.Equal(t, TestableApiAuthToken, result.Token)
}

func Test_Resource_ResourceAbstract_BuildRequestParamsWithAttributes(t *testing.T) {
	config := BuildStubConfig()
	transport := NewHttpTransport(config, nil)
	resource := NewResourceAbstract(transport, config)
	attributes := &RequestParamsAttributes{Id: "123456789", DateTime: BuildStubDateTime()}
	result := resource.buildRequestParams(attributes)
	assert.NotEmpty(t, result)
	assert.Equal(t, attributes.Id, result.Id)
	assert.Equal(t, config.ApiKey, result.Auth.ApiKey)
	assert.Equal(t, TestableApiAuthToken, result.Auth.Token)
}

func Test_Resource_ResourceAbstract_BuildRequestParamsWithoutAttributes(t *testing.T) {
	config := BuildStubConfig()
	transport := NewHttpTransport(config, nil)
	resource := NewResourceAbstract(transport, config)
	result := resource.buildRequestParams(nil)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Id)
	assert.NotEmpty(t, result.Auth.Token)
	assert.Equal(t, config.ApiKey, result.Auth.ApiKey)
}

func Test_Resource_ResourceAbstract_MarshalRequestParams(t *testing.T) {
	config := BuildStubConfig()
	transport := NewHttpTransport(config, nil)
	resource := NewResourceAbstract(transport, config)
	attributes := &RequestParamsAttributes{Id: "123456789", DateTime: BuildStubDateTime()}
	params := resource.buildRequestParams(attributes)
	result, err := resource.marshalRequestParams(params)
	expected := `req=<fasa_request id="123456789"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth></fasa_request>`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))
}

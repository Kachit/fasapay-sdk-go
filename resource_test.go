package fasapay

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ResourceAbstractTestSuite struct {
	suite.Suite
	cfg      *Config
	testable *ResourceAbstract
}

func (suite *ResourceAbstractTestSuite) SetupTest() {
	config := BuildStubConfig()
	transport := NewHttpTransport(config, nil)
	suite.cfg = config
	suite.testable = NewResourceAbstract(transport, config)
}

func (suite *ResourceAbstractTestSuite) TestBuildAuthRequestParams() {
	dt := BuildStubDateTime()
	result := suite.testable.buildAuthRequestParams(dt)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), suite.cfg.ApiKey, result.ApiKey)
	assert.Equal(suite.T(), TestableApiAuthToken, result.Token)
}

func (suite *ResourceAbstractTestSuite) TestBuildRequestParamsWithAttributes() {
	attributes := &RequestParamsAttributes{Id: "123456789", DateTime: BuildStubDateTime()}
	result := suite.testable.buildRequestParams(attributes)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), attributes.Id, result.Id)
	assert.Equal(suite.T(), suite.cfg.ApiKey, result.Auth.ApiKey)
	assert.Equal(suite.T(), TestableApiAuthToken, result.Auth.Token)
}

func (suite *ResourceAbstractTestSuite) TestBuildRequestParamsWithoutAttributes() {
	result := suite.testable.buildRequestParams(nil)
	assert.NotEmpty(suite.T(), result)
	assert.NotEmpty(suite.T(), result.Id)
	assert.NotEmpty(suite.T(), result.Auth.Token)
	assert.Equal(suite.T(), suite.cfg.ApiKey, result.Auth.ApiKey)
}

func (suite *ResourceAbstractTestSuite) TestMarshalRequestParams() {
	attributes := &RequestParamsAttributes{Id: "123456789", DateTime: BuildStubDateTime()}
	params := suite.testable.buildRequestParams(attributes)
	result, err := suite.testable.marshalRequestParams(params)
	expected := `req=<fasa_request id="123456789"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth></fasa_request>`
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(result))
}

func TestResourceAbstractTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceAbstractTestSuite))
}

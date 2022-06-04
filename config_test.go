package fasapay

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
	testable *Config
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.testable = NewConfig("foo", "bar")
}

func (suite *ConfigTestSuite) TestNewConfigByDefault() {
	assert.Equal(suite.T(), ProdAPIUrl, suite.testable.Uri)
	assert.Equal(suite.T(), "foo", suite.testable.ApiKey)
	assert.Equal(suite.T(), "bar", suite.testable.ApiSecretWord)
}

func (suite *ConfigTestSuite) TestNewConfigSandbox() {
	result := NewConfigSandbox("foo", "bar")
	assert.Equal(suite.T(), SandboxAPIUrl, result.Uri)
	assert.Equal(suite.T(), "foo", result.ApiKey)
	assert.Equal(suite.T(), "bar", result.ApiSecretWord)
}

func (suite *ConfigTestSuite) TestIsSandbox() {
	assert.False(suite.T(), suite.testable.IsSandbox())
	suite.testable.Uri = ProdAPIUrlSecond
	assert.False(suite.T(), suite.testable.IsSandbox())
	suite.testable.Uri = SandboxAPIUrl
	assert.True(suite.T(), suite.testable.IsSandbox())
}

func (suite *ConfigTestSuite) TestIsValidSuccess() {
	assert.Nil(suite.T(), suite.testable.IsValid())
	assert.NoError(suite.T(), suite.testable.IsValid())
}

func (suite *ConfigTestSuite) TestIsValidEmptyUri() {
	suite.testable.Uri = ""
	result := suite.testable.IsValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "uri" is empty`, result.Error())
}

func (suite *ConfigTestSuite) TestIsValidEmptyPublicKey() {
	suite.testable.ApiKey = ""
	result := suite.testable.IsValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "api_key" is empty`, result.Error())
}

func (suite *ConfigTestSuite) TestIsValidEmptySecretKey() {
	suite.testable.ApiSecretWord = ""
	result := suite.testable.IsValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "api_secret_word" is empty`, result.Error())
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

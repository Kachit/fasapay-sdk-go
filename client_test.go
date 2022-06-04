package fasapay

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientTestSuite struct {
	suite.Suite
}

func (suite *ClientTestSuite) TestNewClientFromConfigValid() {
	cfg := BuildStubConfig()
	client, err := NewClientFromConfig(cfg, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), client)
}

func (suite *ClientTestSuite) TestNewClientFromConfigInvalid() {
	cfg := BuildStubConfig()
	cfg.Uri = ""
	client, err := NewClientFromConfig(cfg, nil)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), client)
}

func (suite *ClientTestSuite) TestGetAccountsResource() {
	client, _ := NewClientFromConfig(BuildStubConfig(), nil)
	result := client.Accounts()
	assert.NotEmpty(suite.T(), result)
}

func (suite *ClientTestSuite) TestGetTransfersResource() {
	client, _ := NewClientFromConfig(BuildStubConfig(), nil)
	result := client.Transfers()
	assert.NotEmpty(suite.T(), result)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

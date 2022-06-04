package fasapay

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AuthTestSuite struct {
	suite.Suite
}

func (suite *AuthTestSuite) TestGenerateAuthToken() {
	dt := BuildStubDateTime()
	result := generateAuthToken(TestableApiKey, TestableApiSecretWord, dt)
	assert.Equal(suite.T(), TestableApiAuthToken, result)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

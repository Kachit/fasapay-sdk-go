package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Signature_GenerateAuthToken(t *testing.T) {
	dt := BuildStubDateTime()
	result := generateAuthToken(TestableApiKey, TestableApiSecretWord, dt)
	assert.Equal(t, TestableApiAuthToken, result)
}

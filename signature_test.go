package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Signature_GenerateAuthToken(t *testing.T) {
	dt := time.Date(2011, time.Month(7), 20, 15, 30, 0, 0, time.UTC)
	result := generateAuthToken(TestableApiKey, TestableApiSecretWord, dt)
	assert.Equal(t, TestableApiAuthToken, result)
}

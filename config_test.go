package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	result := NewConfig("foo", "bar")
	assert.Equal(t, ProdAPIUrl, result.Uri)
	assert.Equal(t, "foo", result.ApiKey)
	assert.Equal(t, "bar", result.ApiSecretWord)
}

func Test_Config_NewConfigSandbox(t *testing.T) {
	result := NewConfigSandbox("foo", "bar")
	assert.Equal(t, SandboxAPIUrl, result.Uri)
	assert.Equal(t, "foo", result.ApiKey)
	assert.Equal(t, "bar", result.ApiSecretWord)
}

func Test_Config_IsSandbox(t *testing.T) {
	result := NewConfig("foo", "bar")
	assert.False(t, result.IsSandbox())
	result.Uri = ProdAPIUrlSecond
	assert.False(t, result.IsSandbox())
	result.Uri = SandboxAPIUrl
	assert.True(t, result.IsSandbox())
}

func Test_Config_IsValidSuccess(t *testing.T) {
	config := Config{Uri: ProdAPIUrl, ApiKey: "foo", ApiSecretWord: "bar"}
	assert.Nil(t, config.IsValid())
	assert.NoError(t, config.IsValid())
}

func Test_Config_IsValidEmptyUri(t *testing.T) {
	filter := Config{ApiKey: "foo", ApiSecretWord: "bar"}
	result := filter.IsValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "uri" is empty`, result.Error())
}

func Test_Config_IsValidEmptyPublicKey(t *testing.T) {
	filter := Config{Uri: ProdAPIUrl, ApiKey: "", ApiSecretWord: "bar"}
	result := filter.IsValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "api_key" is empty`, result.Error())
}

func Test_Config_IsValidEmptySecretKey(t *testing.T) {
	filter := Config{Uri: ProdAPIUrl, ApiKey: "foo", ApiSecretWord: ""}
	result := filter.IsValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "api_secret_word" is empty`, result.Error())
}

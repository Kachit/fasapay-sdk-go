package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfigValid(t *testing.T) {
	cfg := BuildStubConfig()
	client, err := NewClientFromConfig(cfg, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
}

//func Test_Client_GetAccountsResource(t *testing.T) {
//	client, err := NewClientFromConfig(BuildStubConfig(), nil)
//	result := client.Accounts()
//	assert.NotEmpty(t, result)
//}
//
//func Test_Client_GetTransfersResource(t *testing.T) {
//	client, _ := NewClientFromConfig(BuildStubConfig(), nil)
//	result := client.Transfers()
//	assert.NotEmpty(t, result)
//}

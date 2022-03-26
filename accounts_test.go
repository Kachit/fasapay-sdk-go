package fasapay

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Accounts_GetAccountsResponse_UnmarshalSuccess(t *testing.T) {
	var response GetAccountsResponse
	body, _ := LoadStubResponseData("stubs/accounts/details/success.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1234567", response.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", response.DateTime)
	//accounts
	assert.Equal(t, "Budiman", response.Accounts[0].FullName)
	assert.Equal(t, "FP00001", response.Accounts[0].Account)
	assert.Equal(t, "Store", response.Accounts[0].Status)

	assert.Equal(t, "Ani Permata", response.Accounts[1].FullName)
	assert.Equal(t, "FP00002", response.Accounts[1].Account)
	assert.Equal(t, "Verified", response.Accounts[1].Status)
}

func Test_Accounts_GetAccountsResponse_UnmarshalError(t *testing.T) {
	var response GetAccountsResponse
	body, _ := LoadStubResponseData("stubs/accounts/details/error.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1234567", response.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", response.DateTime)
	//errors
	assert.Equal(t, uint64(41001), response.Errors.Code)
	assert.Equal(t, "account", response.Errors.Mode)
	assert.Equal(t, "", response.Errors.Id)

	assert.Equal(t, uint64(0), response.Errors.Data[0].Code)
	assert.Equal(t, "", response.Errors.Data[0].Attribute)
	assert.Equal(t, "ACCOUNT NOT FOUND", response.Errors.Data[0].Message)
	assert.Equal(t, "FP ACCOUNT FP12345 NOT FOUND", response.Errors.Data[0].Detail)
}

func Test_Accounts_GetBalancesResponse_UnmarshalSuccess(t *testing.T) {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1234567", response.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", response.DateTime)
	//accounts
	assert.Equal(t, 19092587.45, response.Balances[0].IDR)
	assert.Equal(t, 3987.31, response.Balances[0].USD)
}

func Test_Accounts_GetBalancesResponse_UnmarshalError(t *testing.T) {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/error.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1234567", response.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", response.DateTime)
	//errors
	assert.Equal(t, uint64(40901), response.Errors.Code)
	assert.Equal(t, "balance", response.Errors.Mode)
	assert.Equal(t, "", response.Errors.Id)

	assert.Equal(t, uint64(0), response.Errors.Data[0].Code)
	assert.Equal(t, "", response.Errors.Data[0].Attribute)
	assert.Equal(t, "WRONG OR INACTIVE CURRENCY", response.Errors.Data[0].Message)
	assert.Equal(t, "WRONG OR INACTIVE CURRENCY CHY", response.Errors.Data[0].Detail)
}

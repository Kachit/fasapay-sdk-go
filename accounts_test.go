package fasapay

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Accounts_GetAccountsRequest_MarshalXmlSuccess(t *testing.T) {
	xmlRequest := &GetAccountsRequest{RequestParams: BuildStubRequest(), Accounts: []string{"FP00001", "FP00002"}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><account>FP00001</account><account>FP00002</account></fasa_request>`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Accounts_GetAccountsResponse_UnmarshalXmlSuccess(t *testing.T) {
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

func Test_Accounts_GetAccountsResponse_UnmarshalXmlError(t *testing.T) {
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

func Test_Accounts_GetAccountsResponse_MarshalJsonSuccess(t *testing.T) {
	var response GetAccountsResponse
	body, _ := LoadStubResponseData("stubs/accounts/details/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","accounts":[{"fullname":"Budiman","account":"FP00001","status":"Store"},{"fullname":"Ani Permata","account":"FP00002","status":"Verified"}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Accounts_GetAccountsResponse_MarshalJsonError(t *testing.T) {
	var response GetAccountsResponse
	body, _ := LoadStubResponseData("stubs/accounts/details/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","errors":{"mode":"account","code":41001,"data":[{"message":"ACCOUNT NOT FOUND","detail":"FP ACCOUNT FP12345 NOT FOUND"}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Accounts_GetBalancesRequest_MarshalXmlSuccess(t *testing.T) {
	xmlRequest := &GetBalancesRequest{RequestParams: BuildStubRequest(), Balances: []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><balance>IDR</balance><balance>USD</balance></fasa_request>`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Accounts_GetBalancesResponse_UnmarshalXmlSuccess(t *testing.T) {
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

func Test_Accounts_GetBalancesResponse_UnmarshalXmlError(t *testing.T) {
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

func Test_Accounts_GetBalancesResponse_MarshalJsonSuccess(t *testing.T) {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","balances":[{"IDR":19092587.45,"USD":3987.31}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Accounts_GetBalancesResponse_MarshalJsonError(t *testing.T) {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","errors":{"mode":"balance","code":40901,"data":[{"message":"WRONG OR INACTIVE CURRENCY","detail":"WRONG OR INACTIVE CURRENCY CHY"}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

package fasapay

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_Accounts_GetAccountsRequest_MarshalXmlSuccess(t *testing.T) {
	xmlRequest := &GetAccountsRequest{RequestParams: BuildStubRequest(), Accounts: []string{"FP00001", "FP00002"}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><account>FP00001</account><account>FP00002</account></fasa_request>`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Accounts_AccountsResource_GetAccountsSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &AccountsResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}

	body, _ := LoadStubResponseData("stubs/accounts/details/success.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	accounts := []string{"FP0000001"}
	result, resp, err := resource.GetAccounts(accounts, ctx, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.True(t, result.IsSuccess())
	assert.Equal(t, "1234567", result.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", result.DateTime)
	//accounts
	assert.Equal(t, "Budiman", result.Accounts[0].FullName)
	assert.Equal(t, "FP00001", result.Accounts[0].Account)
	assert.Equal(t, "Store", result.Accounts[0].Status)

	assert.Equal(t, "Ani Permata", result.Accounts[1].FullName)
	assert.Equal(t, "FP00002", result.Accounts[1].Account)
	assert.Equal(t, "Verified", result.Accounts[1].Status)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_Accounts_AccountsResource_GetAccountsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &AccountsResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}

	body, _ := LoadStubResponseData("stubs/accounts/details/error.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	accounts := []string{"FP0000001"}
	result, resp, err := resource.GetAccounts(accounts, ctx, nil)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.False(t, result.IsSuccess())
	assert.Equal(t, "1234567", result.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", result.DateTime)
	//errors
	assert.Equal(t, uint64(41001), result.Errors.Code)
	assert.Equal(t, "account", result.Errors.Mode)
	assert.Equal(t, "", result.Errors.Id)

	assert.Equal(t, uint64(0), result.Errors.Data[0].Code)
	assert.Equal(t, "", result.Errors.Data[0].Attribute)
	assert.Equal(t, "ACCOUNT NOT FOUND", result.Errors.Data[0].Message)
	assert.Equal(t, "FP ACCOUNT FP12345 NOT FOUND", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "UNEXPECTED ERROR", err.Error())
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

func Test_Accounts_AccountsResource_GetBalancesSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &AccountsResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}

	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	currencies := []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}
	result, resp, err := resource.GetBalances(currencies, ctx, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.True(t, result.IsSuccess())
	assert.Equal(t, "1234567", result.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", result.DateTime)
	//balances
	assert.Equal(t, 19092587.45, result.Balances.IDR)
	assert.Equal(t, 3987.31, result.Balances.USD)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_Accounts_GetBalancesResponse_GetBalancesError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &AccountsResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}

	body, _ := LoadStubResponseData("stubs/accounts/balances/error.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	currencies := []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}
	result, resp, err := resource.GetBalances(currencies, ctx, nil)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.False(t, result.IsSuccess())
	assert.Equal(t, "1234567", result.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", result.DateTime)
	//errors
	assert.Equal(t, uint64(40901), result.Errors.Code)
	assert.Equal(t, "balance", result.Errors.Mode)
	assert.Equal(t, "", result.Errors.Id)

	assert.Equal(t, uint64(0), result.Errors.Data[0].Code)
	assert.Equal(t, "", result.Errors.Data[0].Attribute)
	assert.Equal(t, "WRONG OR INACTIVE CURRENCY", result.Errors.Data[0].Message)
	assert.Equal(t, "WRONG OR INACTIVE CURRENCY CHY", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "UNEXPECTED ERROR", err.Error())
}

func Test_Accounts_GetBalancesResponse_MarshalJsonSuccess(t *testing.T) {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","balances":{"IDR":19092587.45,"USD":3987.31}}`
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

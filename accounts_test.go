package fasapay

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type AccountsTestSuite struct {
	suite.Suite
}

func (suite *AccountsTestSuite) TestGetAccountsRequestMarshalXmlSuccess() {
	xmlRequest := &GetAccountsRequest{RequestParams: BuildStubRequest(), Accounts: []string{"FP00001", "FP00002"}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><account>FP00001</account><account>FP00002</account></fasa_request>`
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *AccountsTestSuite) TestGetAccountsResponseMarshalJsonSuccess() {
	var response GetAccountsResponse
	body, _ := LoadStubResponseData("stubs/accounts/details/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","accounts":[{"fullname":"Budiman","account":"FP00001","status":"Store"},{"fullname":"Ani Permata","account":"FP00002","status":"Verified"}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *AccountsTestSuite) TestGetAccountsResponseMarshalJsonError() {
	var response GetAccountsResponse
	body, _ := LoadStubResponseData("stubs/accounts/details/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","errors":{"mode":"account","code":41001,"data":[{"message":"ACCOUNT NOT FOUND","detail":"FP ACCOUNT FP12345 NOT FOUND"}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *AccountsTestSuite) TestGetBalancesRequestMarshalXmlSuccess() {
	xmlRequest := &GetBalancesRequest{RequestParams: BuildStubRequest(), Balances: []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><balance>IDR</balance><balance>USD</balance></fasa_request>`
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *AccountsTestSuite) TestGetBalancesResponseMarshalJsonSuccess() {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","balances":{"IDR":19092587.45,"USD":3987.31}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *AccountsTestSuite) TestGetBalancesResponseMarshalJsonError() {
	var response GetBalancesResponse
	body, _ := LoadStubResponseData("stubs/accounts/balances/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","errors":{"mode":"balance","code":40901,"data":[{"message":"WRONG OR INACTIVE CURRENCY","detail":"WRONG OR INACTIVE CURRENCY CHY"}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func TestAccountsTestSuite(t *testing.T) {
	suite.Run(t, new(AccountsTestSuite))
}

type AccountsResourceTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *AccountsResource
}

func (suite *AccountsResourceTestSuite) SetupTest() {
	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()
	suite.cfg = cfg
	suite.ctx = context.Background()
	suite.testable = &AccountsResource{NewResourceAbstract(transport, cfg)}
	httpmock.Activate()
}

func (suite *AccountsResourceTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *AccountsResourceTestSuite) TestGetAccountsSuccess() {
	body, _ := LoadStubResponseData("stubs/accounts/details/success.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	accounts := []string{"FP0000001"}
	result, resp, err := suite.testable.GetAccounts(accounts, suite.ctx, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.True(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1234567", result.Id)
	assert.Equal(suite.T(), "2013-01-01T10:58:43+07:00", result.DateTime)
	//accounts
	assert.Equal(suite.T(), "Budiman", result.Accounts[0].FullName)
	assert.Equal(suite.T(), "FP00001", result.Accounts[0].Account)
	assert.Equal(suite.T(), "Store", result.Accounts[0].Status)

	assert.Equal(suite.T(), "Ani Permata", result.Accounts[1].FullName)
	assert.Equal(suite.T(), "FP00002", result.Accounts[1].Account)
	assert.Equal(suite.T(), "Verified", result.Accounts[1].Status)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *AccountsResourceTestSuite) TestGetAccountsXmlError() {
	body, _ := LoadStubResponseData("stubs/accounts/details/error.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	accounts := []string{"FP0000001"}
	result, resp, err := suite.testable.GetAccounts(accounts, suite.ctx, nil)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1234567", result.Id)
	assert.Equal(suite.T(), "2013-01-01T10:58:43+07:00", result.DateTime)
	//errors
	assert.Equal(suite.T(), uint64(41001), result.Errors.Code)
	assert.Equal(suite.T(), "account", result.Errors.Mode)
	assert.Equal(suite.T(), "", result.Errors.Id)

	assert.Equal(suite.T(), uint64(0), result.Errors.Data[0].Code)
	assert.Equal(suite.T(), "", result.Errors.Data[0].Attribute)
	assert.Equal(suite.T(), "ACCOUNT NOT FOUND", result.Errors.Data[0].Message)
	assert.Equal(suite.T(), "FP ACCOUNT FP12345 NOT FOUND", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "UNEXPECTED ERROR", err.Error())
}

func (suite *AccountsResourceTestSuite) TestGetAccountsNonXmlError() {
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	accounts := []string{"FP0000001"}
	result, resp, err := suite.testable.GetAccounts(accounts, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "AccountsResource.GetAccounts error: EOF", err.Error())
}

func (suite *AccountsResourceTestSuite) TestGetBalancesSuccess() {
	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	currencies := []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}
	result, resp, err := suite.testable.GetBalances(currencies, suite.ctx, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.True(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1234567", result.Id)
	assert.Equal(suite.T(), "2013-01-01T10:58:43+07:00", result.DateTime)
	//balances
	assert.Equal(suite.T(), 19092587.45, result.Balances.IDR)
	assert.Equal(suite.T(), 3987.31, result.Balances.USD)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *AccountsResourceTestSuite) TestGetBalancesXmlError() {
	body, _ := LoadStubResponseData("stubs/accounts/balances/error.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	currencies := []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}
	result, resp, err := suite.testable.GetBalances(currencies, suite.ctx, nil)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1234567", result.Id)
	assert.Equal(suite.T(), "2013-01-01T10:58:43+07:00", result.DateTime)
	//errors
	assert.Equal(suite.T(), uint64(40901), result.Errors.Code)
	assert.Equal(suite.T(), "balance", result.Errors.Mode)
	assert.Equal(suite.T(), "", result.Errors.Id)

	assert.Equal(suite.T(), uint64(0), result.Errors.Data[0].Code)
	assert.Equal(suite.T(), "", result.Errors.Data[0].Attribute)
	assert.Equal(suite.T(), "WRONG OR INACTIVE CURRENCY", result.Errors.Data[0].Message)
	assert.Equal(suite.T(), "WRONG OR INACTIVE CURRENCY CHY", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "UNEXPECTED ERROR", err.Error())
}

func (suite *AccountsResourceTestSuite) TestGetBalancesNonXmlError() {
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	currencies := []CurrencyCode{CurrencyCodeIDR, CurrencyCodeUSD}
	result, resp, err := suite.testable.GetBalances(currencies, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "AccountsResource.GetBalances error: EOF", err.Error())
}

func TestAccountsResourceTestSuite(t *testing.T) {
	suite.Run(t, new(AccountsResourceTestSuite))
}

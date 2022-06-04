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

type TransfersTestSuite struct {
	suite.Suite
}

func (suite *TransfersTestSuite) TestGetHistoryRequestMarshalXmlSuccess() {
	xmlRequest := &GetHistoryRequest{RequestParams: BuildStubRequest(), History: &GetHistoryRequestParams{
		StartDate: "2011-07-01",
		EndDate:   "2011-07-09",
		Type:      "transfer",
		OrderBy:   "date",
		Order:     "DESC",
		Page:      uint64(3),
		PageSize:  uint64(5),
	}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><history><start_date>2011-07-01</start_date><end_date>2011-07-09</end_date><type>transfer</type><order_by>date</order_by><order>DESC</order><page>3</page><page_size>5</page_size></history></fasa_request>`
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *TransfersTestSuite) TestGetHistoryResponseMarshalJsonSuccess() {
	var response GetHistoryResponse
	body, _ := LoadStubResponseData("stubs/transfers/history/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1312342474","date_time":"2011-08-03T10:34:34+07:00","history":{"page":{"total_item":579,"page_count":58,"current_page":0},"details":[{"batchnumber":"TR2011072685119","datetime":"2011-07-26 15:44:35","type":"Keluar","to":"FP10500","from":"FP12049","amount":11160,"note":"Pembayaran untuk pembelian Liberty Reserve","status":"FINISH","currency":"","fee":0},{"batchnumber":"TR2011072521135","datetime":"2011-07-25 11:38:43","type":"Keluar","to":"FP89680","from":"FP12049","amount":1000,"note":"standart operation","status":"FINISH","currency":"","fee":0}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *TransfersTestSuite) TestGetDetailsRequestMarshalXmlSuccess() {
	var detailParam1 GetDetailsRequestDetailParamsString = "foo"
	detailParam2 := &GetDetailsRequestDetailParamsStruct{Ref: "foo"}
	xmlRequest := &GetDetailsRequest{RequestParams: BuildStubRequest(), Details: []GetDetailsDetailParamsInterface{&detailParam1, detailParam2}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><detail>foo</detail><detail><ref>foo</ref></detail></fasa_request>`
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
	assert.Equal(suite.T(), "struct", detailParam2.GetDetailType())
	assert.Equal(suite.T(), "string", detailParam1.GetDetailType())
}

func (suite *TransfersTestSuite) TestGetDetailsResponseMarshalJsonSuccess() {
	var response GetDetailsResponse
	body, _ := LoadStubResponseData("stubs/transfers/details/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","details":[{"mode":"detail","code":210,"batchnumber":"TR2012092791234","date":"2012-10-20","time":"10:09:36","from":"FP00001","to":"FP00002","amount":1000,"total":1100,"currency":"IDR","note":"Payment for something","status":"FINISH","fee":100,"type":"Transfer Out","method":"api_xml","fee_mod":"FiS"}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *TransfersTestSuite) TestGetDetailsResponseMarshalJsonError() {
	var response GetDetailsResponse
	body, _ := LoadStubResponseData("stubs/transfers/details/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","errors":{"mode":"detail","code":40701,"data":[{"message":"TRANSACTION NOT FOUND","detail":"BATCHNUMBER TR2012100291308 NOT FOUND"}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *TransfersTestSuite) TestCreateTransferRequestMarshalXmlSuccess() {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	xmlRequest := &CreateTransferRequest{RequestParams: BuildStubRequest(), Transfers: []*CreateTransferRequestParams{transfer}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><transfer id="123"><to>FP89680</to><amount>1000</amount><currency>IDR</currency><note>standart operation</note></transfer></fasa_request>`
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *TransfersTestSuite) TestCreateTransferRequestIsValidSuccess() {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	assert.Nil(suite.T(), transfer.isValid())
	assert.NoError(suite.T(), transfer.isValid())
}

func (suite *TransfersTestSuite) TestCreateTransferRequestIsValidEmptyParameterTo() {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result := transfer.isValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "to" is empty`, result.Error())
}

func (suite *TransfersTestSuite) TestCreateTransferRequestIsValidEmptyParameterCurrency() {
	transfer := &CreateTransferRequestParams{
		Id:     "123",
		To:     "FP89680",
		Amount: 1000.0,
		Note:   "standart operation",
	}
	result := transfer.isValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "currency" is empty`, result.Error())
}

func (suite *TransfersTestSuite) TestCreateTransferRequestIsValidEmptyParameterAmount() {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result := transfer.isValid()
	assert.Error(suite.T(), result)
	assert.Equal(suite.T(), `parameter "amount" is empty`, result.Error())
}

func (suite *TransfersTestSuite) TestCreateTransferResponseMarshalJsonSuccess() {
	var response CreateTransferResponse
	body, _ := LoadStubResponseData("stubs/transfers/transfer/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1311059195","date_time":"2011-07-19T14:06:35+07:00","transfers":[{"mode":"transfer","code":203,"batchnumber":"TR2011071917277","date":"2011-07-19","time":"14:06:35","from":"FP12049","to":"FP89680","fee":100,"amount":1000,"total":1100,"fee_mode":"FiS","currency":"IDR","note":"standart operation","status":"FINISH","type":"Keluar","balance":2815832,"method":"xml_api"}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func (suite *TransfersTestSuite) TestCreateTransferResponseMarshalJsonError() {
	var response CreateTransferResponse
	body, _ := LoadStubResponseData("stubs/transfers/transfer/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(suite.T(), err)
	expected := `{"id":"1311059195","date_time":"2011-07-19T14:06:35+07:00","errors":{"id":"tid3","mode":"transfer","code":40600,"data":[{"code":40605,"attribute":"id_kurensi","message":"Kurensi tidak boleh kosong."},{"code":40601,"attribute":"to","message":"Tidak ada User dengan Nomor Akun FP89681"},{"code":40602,"attribute":"jumlah","message":"Jumlah melebihi batas yg diijinkan."}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, string(bytes))
}

func TestTransfersTestSuite(t *testing.T) {
	suite.Run(t, new(TransfersTestSuite))
}

type TransfersResourceTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *TransfersResource
}

func (suite *TransfersResourceTestSuite) SetupTest() {
	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()
	suite.cfg = cfg
	suite.ctx = context.Background()
	suite.testable = &TransfersResource{NewResourceAbstract(transport, cfg)}
	httpmock.Activate()
}

func (suite *TransfersResourceTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *TransfersResourceTestSuite) TestGetHistorySuccess() {
	body, _ := LoadStubResponseData("stubs/transfers/history/success.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	historyFilter := &GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
	result, resp, err := suite.testable.GetHistory(historyFilter, suite.ctx, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.True(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1312342474", result.Id)
	assert.Equal(suite.T(), "2011-08-03T10:34:34+07:00", result.DateTime)
	//pagination
	assert.Equal(suite.T(), uint64(579), result.History.Page.TotalItem)
	assert.Equal(suite.T(), uint64(58), result.History.Page.PageCount)
	assert.Equal(suite.T(), uint64(0), result.History.Page.CurrentPage)
	//details
	assert.Equal(suite.T(), "TR2011072685119", result.History.Details[0].BatchNumber)
	assert.Equal(suite.T(), "2011-07-26 15:44:35", result.History.Details[0].Datetime)
	assert.Equal(suite.T(), "Keluar", result.History.Details[0].Type)
	assert.Equal(suite.T(), "FP10500", result.History.Details[0].To)
	assert.Equal(suite.T(), "FP12049", result.History.Details[0].From)
	assert.Equal(suite.T(), 11160.000, result.History.Details[0].Amount)
	assert.Equal(suite.T(), "Pembayaran untuk pembelian Liberty Reserve", result.History.Details[0].Note)
	assert.Equal(suite.T(), "FINISH", result.History.Details[0].Status)

	assert.Equal(suite.T(), "TR2011072521135", result.History.Details[1].BatchNumber)
	assert.Equal(suite.T(), "2011-07-25 11:38:43", result.History.Details[1].Datetime)
	assert.Equal(suite.T(), "Keluar", result.History.Details[1].Type)
	assert.Equal(suite.T(), "FP89680", result.History.Details[1].To)
	assert.Equal(suite.T(), "FP12049", result.History.Details[1].From)
	assert.Equal(suite.T(), 1000.000, result.History.Details[1].Amount)
	assert.Equal(suite.T(), "standart operation", result.History.Details[1].Note)
	assert.Equal(suite.T(), "FINISH", result.History.Details[1].Status)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *TransfersResourceTestSuite) TestGetHistoryXmlError() {
	body, _ := LoadStubResponseData("stubs/transfers/history/error.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	historyFilter := &GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
	result, resp, err := suite.testable.GetHistory(historyFilter, suite.ctx, nil)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1312342474", result.Id)
	assert.Equal(suite.T(), "2011-08-03T10:34:34+07:00", result.DateTime)
	//errors
	assert.Equal(suite.T(), uint64(40701), result.Errors.Code)
	assert.Equal(suite.T(), "history", result.Errors.Mode)
	assert.Equal(suite.T(), "", result.Errors.Id)

	assert.Equal(suite.T(), uint64(0), result.Errors.Data[0].Code)
	assert.Equal(suite.T(), "", result.Errors.Data[0].Attribute)
	assert.Equal(suite.T(), "INVALID DATE FORMAT (yyyy-mm-dd)", result.Errors.Data[0].Message)
	assert.Equal(suite.T(), "INVALID DATE FORMAT (yyyy-mm-dd) foo", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "UNEXPECTED ERROR", err.Error())
}

func (suite *TransfersResourceTestSuite) TestGetHistoryNonXmlError() {
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	historyFilter := &GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
	result, resp, err := suite.testable.GetHistory(historyFilter, suite.ctx, nil)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "TransfersResource.GetHistory error: EOF", err.Error())
}

func (suite *TransfersResourceTestSuite) TestGetDetailsSuccess() {
	body, _ := LoadStubResponseData("stubs/transfers/details/success.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	var detail GetDetailsRequestDetailParamsString = "TR0000000001"
	details := []GetDetailsDetailParamsInterface{&detail}
	result, resp, err := suite.testable.GetDetails(details, suite.ctx, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.True(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1234567", result.Id)
	assert.Equal(suite.T(), "2013-01-01T10:58:43+07:00", result.DateTime)
	//detail
	assert.Equal(suite.T(), "detail", result.Details[0].Mode)
	assert.Equal(suite.T(), uint64(210), result.Details[0].Code)
	assert.Equal(suite.T(), "TR2012092791234", result.Details[0].BatchNumber)
	assert.Equal(suite.T(), "2012-10-20", result.Details[0].Date)
	assert.Equal(suite.T(), "10:09:36", result.Details[0].Time)
	assert.Equal(suite.T(), "FP00001", result.Details[0].From)
	assert.Equal(suite.T(), "FP00002", result.Details[0].To)
	assert.Equal(suite.T(), 1000.000, result.Details[0].Amount)
	assert.Equal(suite.T(), 100.000, result.Details[0].Fee)
	assert.Equal(suite.T(), float64(1100), result.Details[0].Total)
	assert.Equal(suite.T(), "FiS", result.Details[0].FeeMode)
	assert.Equal(suite.T(), "IDR", result.Details[0].Currency)
	assert.Equal(suite.T(), "Payment for something", result.Details[0].Note)
	assert.Equal(suite.T(), "FINISH", result.Details[0].Status)
	assert.Equal(suite.T(), "Transfer Out", result.Details[0].Type)
	assert.Equal(suite.T(), "api_xml", result.Details[0].Method)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *TransfersResourceTestSuite) TestGetDetailsXmlError() {
	body, _ := LoadStubResponseData("stubs/transfers/details/error.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	var detail GetDetailsRequestDetailParamsString = "TR0000000001"
	details := []GetDetailsDetailParamsInterface{&detail}
	result, resp, err := suite.testable.GetDetails(details, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1234567", result.Id)
	assert.Equal(suite.T(), "2013-01-01T10:58:43+07:00", result.DateTime)
	//errors
	assert.Equal(suite.T(), uint64(40701), result.Errors.Code)
	assert.Equal(suite.T(), "detail", result.Errors.Mode)
	assert.Equal(suite.T(), "", result.Errors.Id)

	assert.Equal(suite.T(), uint64(0), result.Errors.Data[0].Code)
	assert.Equal(suite.T(), "", result.Errors.Data[0].Attribute)
	assert.Equal(suite.T(), "TRANSACTION NOT FOUND", result.Errors.Data[0].Message)
	assert.Equal(suite.T(), "BATCHNUMBER TR2012100291308 NOT FOUND", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "UNEXPECTED ERROR", err.Error())
}

func (suite *TransfersResourceTestSuite) TestGetDetailsNonXmlError() {
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	var detail GetDetailsRequestDetailParamsString = "TR0000000001"
	details := []GetDetailsDetailParamsInterface{&detail}
	result, resp, err := suite.testable.GetDetails(details, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "TransfersResource.GetDetails error: EOF", err.Error())
}

func (suite *TransfersResourceTestSuite) TestCreateTransferSuccess() {
	body, _ := LoadStubResponseData("stubs/transfers/transfer/success.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := suite.testable.CreateTransfer([]*CreateTransferRequestParams{transfer}, suite.ctx, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.True(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1311059195", result.Id)
	assert.Equal(suite.T(), "2011-07-19T14:06:35+07:00", result.DateTime)
	//transfer
	assert.Equal(suite.T(), "transfer", result.Transfers[0].Mode)
	assert.Equal(suite.T(), uint64(203), result.Transfers[0].Code)
	assert.Equal(suite.T(), "TR2011071917277", result.Transfers[0].BatchNumber)
	assert.Equal(suite.T(), "2011-07-19", result.Transfers[0].Date)
	assert.Equal(suite.T(), "14:06:35", result.Transfers[0].Time)
	assert.Equal(suite.T(), "FP12049", result.Transfers[0].From)
	assert.Equal(suite.T(), "FP89680", result.Transfers[0].To)
	assert.Equal(suite.T(), 1000.0, result.Transfers[0].Amount)
	assert.Equal(suite.T(), float64(100), result.Transfers[0].Fee)
	assert.Equal(suite.T(), 1100.0, result.Transfers[0].Total)
	assert.Equal(suite.T(), "FiS", result.Transfers[0].FeeMode)
	assert.Equal(suite.T(), "IDR", result.Transfers[0].Currency)
	assert.Equal(suite.T(), "standart operation", result.Transfers[0].Note)
	assert.Equal(suite.T(), "FINISH", result.Transfers[0].Status)
	assert.Equal(suite.T(), "Keluar", result.Transfers[0].Type)
	assert.Equal(suite.T(), 2815832.00, result.Transfers[0].Balance)
	assert.Equal(suite.T(), "xml_api", result.Transfers[0].Method)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *TransfersResourceTestSuite) TestCreateTransferRequestError() {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := suite.testable.CreateTransfer([]*CreateTransferRequestParams{transfer}, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	//error
	assert.Equal(suite.T(), `TransfersResource.CreateTransfer error: parameter "to" is empty`, err.Error())
}

func (suite *TransfersResourceTestSuite) TestCreateTransferXmlError() {
	body, _ := LoadStubResponseData("stubs/transfers/transfer/error.xml")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := suite.testable.CreateTransfer([]*CreateTransferRequestParams{transfer}, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	//common
	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "1311059195", result.Id)
	assert.Equal(suite.T(), "2011-07-19T14:06:35+07:00", result.DateTime)
	//errors
	assert.Equal(suite.T(), uint64(40600), result.Errors.Code)
	assert.Equal(suite.T(), "transfer", result.Errors.Mode)
	assert.Equal(suite.T(), "tid3", result.Errors.Id)

	assert.Equal(suite.T(), uint64(40605), result.Errors.Data[0].Code)
	assert.Equal(suite.T(), "id_kurensi", result.Errors.Data[0].Attribute)
	assert.Equal(suite.T(), "Kurensi tidak boleh kosong.", result.Errors.Data[0].Message)
	assert.Equal(suite.T(), "", result.Errors.Data[0].Detail)

	assert.Equal(suite.T(), uint64(40601), result.Errors.Data[1].Code)
	assert.Equal(suite.T(), "to", result.Errors.Data[1].Attribute)
	assert.Equal(suite.T(), "Tidak ada User dengan Nomor Akun FP89681", result.Errors.Data[1].Message)
	assert.Equal(suite.T(), "", result.Errors.Data[1].Detail)

	assert.Equal(suite.T(), uint64(40602), result.Errors.Data[2].Code)
	assert.Equal(suite.T(), "jumlah", result.Errors.Data[2].Attribute)
	assert.Equal(suite.T(), "Jumlah melebihi batas yg diijinkan.", result.Errors.Data[2].Message)
	assert.Equal(suite.T(), "", result.Errors.Data[2].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "NOT ACCEPTABLE TRANSFER", err.Error())
}

func (suite *TransfersResourceTestSuite) TestCreateTransferNonXmlError() {
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := suite.testable.CreateTransfer([]*CreateTransferRequestParams{transfer}, suite.ctx, nil)

	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
	//error
	assert.Equal(suite.T(), "TransfersResource.CreateTransfer error: EOF", err.Error())
}

func (suite *TransfersResourceTestSuite) TestValidateTransferParamsValid() {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	err := suite.testable.validateTransferParams([]*CreateTransferRequestParams{transfer})
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TransfersResourceTestSuite) TestValidateTransferParamsInvalid() {
	transfer1 := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	transfer2 := &CreateTransferRequestParams{
		Id:       "123",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	err := suite.testable.validateTransferParams([]*CreateTransferRequestParams{transfer1, transfer2})
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), `parameter "to" is empty`, err.Error())
}

func TestTransfersResourceTestSuite(t *testing.T) {
	suite.Run(t, new(TransfersResourceTestSuite))
}

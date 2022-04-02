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

func Test_Transfers_GetHistoryRequest_MarshalXmlSuccess(t *testing.T) {
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
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Transfers_TransfersResource_GetHistorySuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/transfers/history/success.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	historyFilter := &GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
	result, resp, err := resource.GetHistory(historyFilter, ctx, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.True(t, result.IsSuccess())
	assert.Equal(t, "1312342474", result.Id)
	assert.Equal(t, "2011-08-03T10:34:34+07:00", result.DateTime)
	//pagination
	assert.Equal(t, uint64(579), result.History.Page.TotalItem)
	assert.Equal(t, uint64(58), result.History.Page.PageCount)
	assert.Equal(t, uint64(0), result.History.Page.CurrentPage)
	//details
	assert.Equal(t, "TR2011072685119", result.History.Details[0].BatchNumber)
	assert.Equal(t, "2011-07-26 15:44:35", result.History.Details[0].Datetime)
	assert.Equal(t, "Keluar", result.History.Details[0].Type)
	assert.Equal(t, "FP10500", result.History.Details[0].To)
	assert.Equal(t, "FP12049", result.History.Details[0].From)
	assert.Equal(t, 11160.000, result.History.Details[0].Amount)
	assert.Equal(t, "Pembayaran untuk pembelian Liberty Reserve", result.History.Details[0].Note)
	assert.Equal(t, "FINISH", result.History.Details[0].Status)

	assert.Equal(t, "TR2011072521135", result.History.Details[1].BatchNumber)
	assert.Equal(t, "2011-07-25 11:38:43", result.History.Details[1].Datetime)
	assert.Equal(t, "Keluar", result.History.Details[1].Type)
	assert.Equal(t, "FP89680", result.History.Details[1].To)
	assert.Equal(t, "FP12049", result.History.Details[1].From)
	assert.Equal(t, 1000.000, result.History.Details[1].Amount)
	assert.Equal(t, "standart operation", result.History.Details[1].Note)
	assert.Equal(t, "FINISH", result.History.Details[1].Status)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_Transfers_TransfersResource_GetHistoryXmlError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/transfers/history/error.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	historyFilter := &GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
	result, resp, err := resource.GetHistory(historyFilter, ctx, nil)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.False(t, result.IsSuccess())
	assert.Equal(t, "1312342474", result.Id)
	assert.Equal(t, "2011-08-03T10:34:34+07:00", result.DateTime)
	//errors
	assert.Equal(t, uint64(40701), result.Errors.Code)
	assert.Equal(t, "history", result.Errors.Mode)
	assert.Equal(t, "", result.Errors.Id)

	assert.Equal(t, uint64(0), result.Errors.Data[0].Code)
	assert.Equal(t, "", result.Errors.Data[0].Attribute)
	assert.Equal(t, "INVALID DATE FORMAT (yyyy-mm-dd)", result.Errors.Data[0].Message)
	assert.Equal(t, "INVALID DATE FORMAT (yyyy-mm-dd) foo", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "UNEXPECTED ERROR", err.Error())
}

func Test_Transfers_TransfersResource_GetHistoryNonXmlError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	historyFilter := &GetHistoryRequestParams{StartDate: "2022-03-01", EndDate: "2022-03-28"}
	result, resp, err := resource.GetHistory(historyFilter, ctx, nil)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.Empty(t, result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "TransfersResource.GetHistory error: EOF", err.Error())
}

func Test_Transfers_GetHistoryResponse_MarshalJsonSuccess(t *testing.T) {
	var response GetHistoryResponse
	body, _ := LoadStubResponseData("stubs/transfers/history/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1312342474","date_time":"2011-08-03T10:34:34+07:00","history":{"page":{"total_item":579,"page_count":58,"current_page":0},"details":[{"batchnumber":"TR2011072685119","datetime":"2011-07-26 15:44:35","type":"Keluar","to":"FP10500","from":"FP12049","amount":11160,"note":"Pembayaran untuk pembelian Liberty Reserve","status":"FINISH","currency":"","fee":0},{"batchnumber":"TR2011072521135","datetime":"2011-07-25 11:38:43","type":"Keluar","to":"FP89680","from":"FP12049","amount":1000,"note":"standart operation","status":"FINISH","currency":"","fee":0}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Transfers_GetDetailsRequest_MarshalXmlSuccess(t *testing.T) {
	var detailParam1 GetDetailsRequestDetailParamsString = "foo"
	detailParam2 := &GetDetailsRequestDetailParamsStruct{Ref: "foo"}
	xmlRequest := &GetDetailsRequest{RequestParams: BuildStubRequest(), Details: []GetDetailsDetailParamsInterface{&detailParam1, detailParam2}}
	bytes, err := xml.Marshal(xmlRequest)
	expected := `<fasa_request id="1234567"><auth><api_key>11123548cd3a5e5613325132112becf</api_key><token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token></auth><detail>foo</detail><detail><ref>foo</ref></detail></fasa_request>`
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
	assert.Equal(t, "struct", detailParam2.GetDetailType())
	assert.Equal(t, "string", detailParam1.GetDetailType())
}

func Test_Transfers_TransfersResource_GetDetailsSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/transfers/details/success.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	var detail GetDetailsRequestDetailParamsString = "TR0000000001"
	details := []GetDetailsDetailParamsInterface{&detail}
	result, resp, err := resource.GetDetails(details, ctx, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.True(t, result.IsSuccess())
	assert.Equal(t, "1234567", result.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", result.DateTime)
	//detail
	assert.Equal(t, "detail", result.Details[0].Mode)
	assert.Equal(t, uint64(210), result.Details[0].Code)
	assert.Equal(t, "TR2012092791234", result.Details[0].BatchNumber)
	assert.Equal(t, "2012-10-20", result.Details[0].Date)
	assert.Equal(t, "10:09:36", result.Details[0].Time)
	assert.Equal(t, "FP00001", result.Details[0].From)
	assert.Equal(t, "FP00002", result.Details[0].To)
	assert.Equal(t, 1000.000, result.Details[0].Amount)
	assert.Equal(t, 100.000, result.Details[0].Fee)
	assert.Equal(t, float64(1100), result.Details[0].Total)
	assert.Equal(t, "FiS", result.Details[0].FeeMode)
	assert.Equal(t, "IDR", result.Details[0].Currency)
	assert.Equal(t, "Payment for something", result.Details[0].Note)
	assert.Equal(t, "FINISH", result.Details[0].Status)
	assert.Equal(t, "Transfer Out", result.Details[0].Type)
	assert.Equal(t, "api_xml", result.Details[0].Method)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_Transfers_TransfersResource_GetDetailsXmlError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/transfers/details/error.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	var detail GetDetailsRequestDetailParamsString = "TR0000000001"
	details := []GetDetailsDetailParamsInterface{&detail}
	result, resp, err := resource.GetDetails(details, ctx, nil)

	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.False(t, result.IsSuccess())
	assert.Equal(t, "1234567", result.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", result.DateTime)
	//errors
	assert.Equal(t, uint64(40701), result.Errors.Code)
	assert.Equal(t, "detail", result.Errors.Mode)
	assert.Equal(t, "", result.Errors.Id)

	assert.Equal(t, uint64(0), result.Errors.Data[0].Code)
	assert.Equal(t, "", result.Errors.Data[0].Attribute)
	assert.Equal(t, "TRANSACTION NOT FOUND", result.Errors.Data[0].Message)
	assert.Equal(t, "BATCHNUMBER TR2012100291308 NOT FOUND", result.Errors.Data[0].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "UNEXPECTED ERROR", err.Error())
}

func Test_Transfers_TransfersResource_GetDetailsNonXmlError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	var detail GetDetailsRequestDetailParamsString = "TR0000000001"
	details := []GetDetailsDetailParamsInterface{&detail}
	result, resp, err := resource.GetDetails(details, ctx, nil)

	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.Empty(t, result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "TransfersResource.GetDetails error: EOF", err.Error())
}

func Test_Transfers_GetDetailsResponse_MarshalJsonSuccess(t *testing.T) {
	var response GetDetailsResponse
	body, _ := LoadStubResponseData("stubs/transfers/details/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","details":[{"mode":"detail","code":210,"batchnumber":"TR2012092791234","date":"2012-10-20","time":"10:09:36","from":"FP00001","to":"FP00002","amount":1000,"total":1100,"currency":"IDR","note":"Payment for something","status":"FINISH","fee":100,"type":"Transfer Out","method":"api_xml","fee_mod":"FiS"}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Transfers_GetDetailsResponse_MarshalJsonError(t *testing.T) {
	var response GetDetailsResponse
	body, _ := LoadStubResponseData("stubs/transfers/details/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1234567","date_time":"2013-01-01T10:58:43+07:00","errors":{"mode":"detail","code":40701,"data":[{"message":"TRANSACTION NOT FOUND","detail":"BATCHNUMBER TR2012100291308 NOT FOUND"}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Transfers_CreateTransferRequest_MarshalXmlSuccess(t *testing.T) {
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
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Transfers_CreateTransferRequest_IsValidSuccess(t *testing.T) {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	assert.Nil(t, transfer.isValid())
	assert.NoError(t, transfer.isValid())
}

func Test_Transfers_CreateTransferRequest_IsValidEmptyParameterTo(t *testing.T) {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result := transfer.isValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "to" is empty`, result.Error())
}

func Test_Transfers_CreateTransferRequest_IsValidEmptyParameterCurrency(t *testing.T) {
	transfer := &CreateTransferRequestParams{
		Id:     "123",
		To:     "FP89680",
		Amount: 1000.0,
		Note:   "standart operation",
	}
	result := transfer.isValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "currency" is empty`, result.Error())
}

func Test_Transfers_CreateTransferRequest_IsValidEmptyParameterAmount(t *testing.T) {
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result := transfer.isValid()
	assert.Error(t, result)
	assert.Equal(t, `parameter "amount" is empty`, result.Error())
}

func Test_Transfers_TransfersResource_CreateTransferSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/transfers/transfer/success.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := resource.CreateTransfer([]*CreateTransferRequestParams{transfer}, ctx, nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.True(t, result.IsSuccess())
	assert.Equal(t, "1311059195", result.Id)
	assert.Equal(t, "2011-07-19T14:06:35+07:00", result.DateTime)
	//transfer
	assert.Equal(t, "transfer", result.Transfers[0].Mode)
	assert.Equal(t, uint64(203), result.Transfers[0].Code)
	assert.Equal(t, "TR2011071917277", result.Transfers[0].BatchNumber)
	assert.Equal(t, "2011-07-19", result.Transfers[0].Date)
	assert.Equal(t, "14:06:35", result.Transfers[0].Time)
	assert.Equal(t, "FP12049", result.Transfers[0].From)
	assert.Equal(t, "FP89680", result.Transfers[0].To)
	assert.Equal(t, 1000.0, result.Transfers[0].Amount)
	assert.Equal(t, float64(100), result.Transfers[0].Fee)
	assert.Equal(t, 1100.0, result.Transfers[0].Total)
	assert.Equal(t, "FiS", result.Transfers[0].FeeMode)
	assert.Equal(t, "IDR", result.Transfers[0].Currency)
	assert.Equal(t, "standart operation", result.Transfers[0].Note)
	assert.Equal(t, "FINISH", result.Transfers[0].Status)
	assert.Equal(t, "Keluar", result.Transfers[0].Type)
	assert.Equal(t, 2815832.00, result.Transfers[0].Balance)
	assert.Equal(t, "xml_api", result.Transfers[0].Method)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
}

func Test_Transfers_TransfersResource_CreateTransferRequestError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}

	ctx := context.Background()
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := resource.CreateTransfer([]*CreateTransferRequestParams{transfer}, ctx, nil)

	assert.Error(t, err)
	assert.Empty(t, resp)
	assert.Empty(t, result)
	//error
	assert.Equal(t, `TransfersResource.CreateTransfer error: parameter "to" is empty`, err.Error())
}

func Test_Transfers_TransfersResource_CreateTransferXmlError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/transfers/transfer/error.xml")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := resource.CreateTransfer([]*CreateTransferRequestParams{transfer}, ctx, nil)

	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	//common
	assert.False(t, result.IsSuccess())
	assert.Equal(t, "1311059195", result.Id)
	assert.Equal(t, "2011-07-19T14:06:35+07:00", result.DateTime)
	//errors
	assert.Equal(t, uint64(40600), result.Errors.Code)
	assert.Equal(t, "transfer", result.Errors.Mode)
	assert.Equal(t, "tid3", result.Errors.Id)

	assert.Equal(t, uint64(40605), result.Errors.Data[0].Code)
	assert.Equal(t, "id_kurensi", result.Errors.Data[0].Attribute)
	assert.Equal(t, "Kurensi tidak boleh kosong.", result.Errors.Data[0].Message)
	assert.Equal(t, "", result.Errors.Data[0].Detail)

	assert.Equal(t, uint64(40601), result.Errors.Data[1].Code)
	assert.Equal(t, "to", result.Errors.Data[1].Attribute)
	assert.Equal(t, "Tidak ada User dengan Nomor Akun FP89681", result.Errors.Data[1].Message)
	assert.Equal(t, "", result.Errors.Data[1].Detail)

	assert.Equal(t, uint64(40602), result.Errors.Data[2].Code)
	assert.Equal(t, "jumlah", result.Errors.Data[2].Attribute)
	assert.Equal(t, "Jumlah melebihi batas yg diijinkan.", result.Errors.Data[2].Message)
	assert.Equal(t, "", result.Errors.Data[2].Detail)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "NOT ACCEPTABLE TRANSFER", err.Error())
}

func Test_Transfers_TransfersResource_CreateTransferNonXmlError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := BuildStubConfig()
	transport := BuildStubHttpTransport()

	resource := &TransfersResource{ResourceAbstract: NewResourceAbstract(transport, cfg)}
	body, _ := LoadStubResponseData("stubs/errors/500.html")
	httpmock.RegisterResponder(http.MethodPost, cfg.Uri, httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	result, resp, err := resource.CreateTransfer([]*CreateTransferRequestParams{transfer}, ctx, nil)

	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.Empty(t, result)
	//response
	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, body, bodyRsp)
	//error
	assert.Equal(t, "TransfersResource.CreateTransfer error: EOF", err.Error())
}

func Test_Transfers_TransfersResource_ValidateTransferParamsValid(t *testing.T) {
	resource := &TransfersResource{}
	transfer := &CreateTransferRequestParams{
		Id:       "123",
		To:       "FP89680",
		Amount:   1000.0,
		Currency: CurrencyCodeIDR,
		Note:     "standart operation",
	}
	err := resource.validateTransferParams([]*CreateTransferRequestParams{transfer})
	assert.Nil(t, err)
	assert.NoError(t, err)
}

func Test_Transfers_TransfersResource_ValidateTransferParamsInvalid(t *testing.T) {
	resource := &TransfersResource{}
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
	err := resource.validateTransferParams([]*CreateTransferRequestParams{transfer1, transfer2})
	assert.Error(t, err)
	assert.Equal(t, `parameter "to" is empty`, err.Error())
}

func Test_Transfers_CreateTransferResponse_MarshalJsonSuccess(t *testing.T) {
	var response CreateTransferResponse
	body, _ := LoadStubResponseData("stubs/transfers/transfer/success.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1311059195","date_time":"2011-07-19T14:06:35+07:00","transfers":[{"mode":"transfer","code":203,"batchnumber":"TR2011071917277","date":"2011-07-19","time":"14:06:35","from":"FP12049","to":"FP89680","fee":100,"amount":1000,"total":1100,"fee_mode":"FiS","currency":"IDR","note":"standart operation","status":"FINISH","type":"Keluar","balance":2815832,"method":"xml_api"}]}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

func Test_Transfers_CreateTransferResponse_MarshalJsonError(t *testing.T) {
	var response CreateTransferResponse
	body, _ := LoadStubResponseData("stubs/transfers/transfer/error.xml")
	err := xml.Unmarshal(body, &response)

	assert.NoError(t, err)
	expected := `{"id":"1311059195","date_time":"2011-07-19T14:06:35+07:00","errors":{"id":"tid3","mode":"transfer","code":40600,"data":[{"code":40605,"attribute":"id_kurensi","message":"Kurensi tidak boleh kosong."},{"code":40601,"attribute":"to","message":"Tidak ada User dengan Nomor Akun FP89681"},{"code":40602,"attribute":"jumlah","message":"Jumlah melebihi batas yg diijinkan."}]}}`
	bytes, err := json.Marshal(&response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(bytes))
}

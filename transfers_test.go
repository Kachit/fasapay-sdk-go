package fasapay

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
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

func Test_Transfers_GetHistoryResponse_UnmarshalXmlSuccess(t *testing.T) {
	var response GetHistoryResponse
	body, _ := LoadStubResponseData("stubs/transfers/history/success.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1312342474", response.Id)
	assert.Equal(t, "2011-08-03T10:34:34+07:00", response.DateTime)
	//pagination
	assert.Equal(t, uint64(579), response.History.Page.TotalItem)
	assert.Equal(t, uint64(58), response.History.Page.PageCount)
	assert.Equal(t, uint64(0), response.History.Page.CurrentPage)
	//details
	assert.Equal(t, "TR2011072685119", response.History.Details[0].BatchNumber)
	assert.Equal(t, "2011-07-26 15:44:35", response.History.Details[0].Datetime)
	assert.Equal(t, "Keluar", response.History.Details[0].Type)
	assert.Equal(t, "FP10500", response.History.Details[0].To)
	assert.Equal(t, "FP12049", response.History.Details[0].From)
	assert.Equal(t, 11160.000, response.History.Details[0].Amount)
	assert.Equal(t, "Pembayaran untuk pembelian Liberty Reserve", response.History.Details[0].Note)
	assert.Equal(t, "FINISH", response.History.Details[0].Status)

	assert.Equal(t, "TR2011072521135", response.History.Details[1].BatchNumber)
	assert.Equal(t, "2011-07-25 11:38:43", response.History.Details[1].Datetime)
	assert.Equal(t, "Keluar", response.History.Details[1].Type)
	assert.Equal(t, "FP89680", response.History.Details[1].To)
	assert.Equal(t, "FP12049", response.History.Details[1].From)
	assert.Equal(t, 1000.000, response.History.Details[1].Amount)
	assert.Equal(t, "standart operation", response.History.Details[1].Note)
	assert.Equal(t, "FINISH", response.History.Details[1].Status)
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
}

func Test_Transfers_GetDetailsResponse_UnmarshalXmlSuccess(t *testing.T) {
	var response GetDetailsResponse
	body, _ := LoadStubResponseData("stubs/transfers/details/success.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1234567", response.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", response.DateTime)
	//detail
	assert.Equal(t, "detail", response.Details[0].Mode)
	assert.Equal(t, uint64(210), response.Details[0].Code)
	assert.Equal(t, "TR2012092791234", response.Details[0].BatchNumber)
	assert.Equal(t, "2012-10-20", response.Details[0].Date)
	assert.Equal(t, "10:09:36", response.Details[0].Time)
	assert.Equal(t, "FP00001", response.Details[0].From)
	assert.Equal(t, "FP00002", response.Details[0].To)
	assert.Equal(t, 1000.000, response.Details[0].Amount)
	assert.Equal(t, 100.000, response.Details[0].Fee)
	assert.Equal(t, float64(1100), response.Details[0].Total)
	assert.Equal(t, "FiS", response.Details[0].FeeMode)
	assert.Equal(t, "IDR", response.Details[0].Currency)
	assert.Equal(t, "Payment for something", response.Details[0].Note)
	assert.Equal(t, "FINISH", response.Details[0].Status)
	assert.Equal(t, "Transfer Out", response.Details[0].Type)
	assert.Equal(t, "api_xml", response.Details[0].Method)
}

func Test_Transfers_GetDetailsResponse_UnmarshalXmlError(t *testing.T) {
	var response GetDetailsResponse
	body, _ := LoadStubResponseData("stubs/transfers/details/error.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1234567", response.Id)
	assert.Equal(t, "2013-01-01T10:58:43+07:00", response.DateTime)
	//errors
	assert.Equal(t, uint64(40701), response.Errors.Code)
	assert.Equal(t, "detail", response.Errors.Mode)
	assert.Equal(t, "", response.Errors.Id)

	assert.Equal(t, uint64(0), response.Errors.Data[0].Code)
	assert.Equal(t, "", response.Errors.Data[0].Attribute)
	assert.Equal(t, "TRANSACTION NOT FOUND", response.Errors.Data[0].Message)
	assert.Equal(t, "BATCHNUMBER TR2012100291308 NOT FOUND", response.Errors.Data[0].Detail)
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

func Test_Transfers_CreateTransferResponse_UnmarshalXmlSuccess(t *testing.T) {
	var response CreateTransferResponse
	body, _ := LoadStubResponseData("stubs/transfers/transfer/success.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1311059195", response.Id)
	assert.Equal(t, "2011-07-19T14:06:35+07:00", response.DateTime)
	//transfer
	assert.Equal(t, "transfer", response.Transfers[0].Mode)
	assert.Equal(t, uint64(203), response.Transfers[0].Code)
	assert.Equal(t, "TR2011071917277", response.Transfers[0].BatchNumber)
	assert.Equal(t, "2011-07-19", response.Transfers[0].Date)
	assert.Equal(t, "14:06:35", response.Transfers[0].Time)
	assert.Equal(t, "FP12049", response.Transfers[0].From)
	assert.Equal(t, "FP89680", response.Transfers[0].To)
	assert.Equal(t, 1000.0, response.Transfers[0].Amount)
	assert.Equal(t, float64(100), response.Transfers[0].Fee)
	assert.Equal(t, 1100.0, response.Transfers[0].Total)
	assert.Equal(t, "FiS", response.Transfers[0].FeeMode)
	assert.Equal(t, "IDR", response.Transfers[0].Currency)
	assert.Equal(t, "standart operation", response.Transfers[0].Note)
	assert.Equal(t, "FINISH", response.Transfers[0].Status)
	assert.Equal(t, "Keluar", response.Transfers[0].Type)
	assert.Equal(t, 2815832.00, response.Transfers[0].Balance)
	assert.Equal(t, "xml_api", response.Transfers[0].Method)
}

func Test_Transfers_CreateTransferResponse_UnmarshalXmlError(t *testing.T) {
	var response CreateTransferResponse
	body, _ := LoadStubResponseData("stubs/transfers/transfer/error.xml")
	err := xml.Unmarshal(body, &response)
	assert.NoError(t, err)
	//common
	assert.Equal(t, "1311059195", response.Id)
	assert.Equal(t, "2011-07-19T14:06:35+07:00", response.DateTime)
	//errors
	assert.Equal(t, uint64(40600), response.Errors.Code)
	assert.Equal(t, "transfer", response.Errors.Mode)
	assert.Equal(t, "tid3", response.Errors.Id)

	assert.Equal(t, uint64(40605), response.Errors.Data[0].Code)
	assert.Equal(t, "id_kurensi", response.Errors.Data[0].Attribute)
	assert.Equal(t, "Kurensi tidak boleh kosong.", response.Errors.Data[0].Message)
	assert.Equal(t, "", response.Errors.Data[0].Detail)

	assert.Equal(t, uint64(40601), response.Errors.Data[1].Code)
	assert.Equal(t, "to", response.Errors.Data[1].Attribute)
	assert.Equal(t, "Tidak ada User dengan Nomor Akun FP89681", response.Errors.Data[1].Message)
	assert.Equal(t, "", response.Errors.Data[1].Detail)

	assert.Equal(t, uint64(40602), response.Errors.Data[2].Code)
	assert.Equal(t, "jumlah", response.Errors.Data[2].Attribute)
	assert.Equal(t, "Jumlah melebihi batas yg diijinkan.", response.Errors.Data[2].Message)
	assert.Equal(t, "", response.Errors.Data[2].Detail)
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

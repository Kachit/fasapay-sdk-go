package fasapay

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Transfers_GetHistoryResponse_UnmarshalSuccess(t *testing.T) {
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

func Test_Transfers_GetDetailsResponse_UnmarshalSuccess(t *testing.T) {
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

func Test_Transfers_GetDetailsResponse_UnmarshalError(t *testing.T) {
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

func Test_Transfers_CreateTransferResponse_UnmarshalSuccess(t *testing.T) {
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

func Test_Transfers_CreateTransferResponse_UnmarshalError(t *testing.T) {
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

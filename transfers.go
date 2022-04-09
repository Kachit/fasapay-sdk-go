package fasapay

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

//CreateTransferRequestParams struct
type CreateTransferRequestParams struct {
	XMLName  xml.Name           `xml:"transfer"`
	Id       string             `xml:"id,attr,omitempty" json:"id"`        //id transfer for marking the transfer (max 50 character)
	To       string             `xml:"to" json:"to"`                       //is the FasaPay account target format: FPnnnnn
	Amount   float64            `xml:"amount" json:"amount"`               //is the amount of the transferred fund. with point (.) as the decimal separator
	Currency CurrencyCode       `xml:"currency" json:"currency"`           //is the currency used in the transfer (IDR | USD)
	FeeMode  TransactionFeeMode `xml:"fee_mode,omitempty" json:"fee_mode"` //is Fee Mode used in the transfer. default to FiR (FiR | FiS)
	Note     string             `xml:"note,omitempty" json:"note"`         //is note of the transfer (max 255 character)
	Ref      string             `xml:"ref,omitempty" json:"ref"`           //Reference Code that can be used to track transaction (max 50 character)
}

//isValid method
func (ctr *CreateTransferRequestParams) isValid() error {
	var err error
	if ctr.To == "" {
		err = fmt.Errorf(`parameter "to" is empty`)
	} else if ctr.Currency == "" {
		err = fmt.Errorf(`parameter "currency" is empty`)
	} else if ctr.Amount == 0 {
		err = fmt.Errorf(`parameter "amount" is empty`)
	}
	return err
}

//CreateTransferRequest struct
type CreateTransferRequest struct {
	*RequestParams
	Transfers []*CreateTransferRequestParams `xml:"transfer" json:"transfers"`
}

//CreateTransferResponse struct
type CreateTransferResponse struct {
	*ResponseBody
	Transfers []*CreateTransferResponseParams `xml:"transfer,omitempty" json:"transfers,omitempty"`
}

//CreateTransferResponseParams struct
type CreateTransferResponseParams struct {
	Mode        string  `xml:"mode,attr" json:"mode"`
	Code        uint64  `xml:"code,attr" json:"code"`
	BatchNumber string  `xml:"batchnumber" json:"batchnumber"`
	Date        string  `xml:"date" json:"date"`
	Time        string  `xml:"time" json:"time"`
	From        string  `xml:"from" json:"from"`
	To          string  `xml:"to" json:"to"`
	Fee         float64 `xml:"fee" json:"fee"`
	Amount      float64 `xml:"amount" json:"amount"`
	Total       float64 `xml:"total" json:"total"`
	FeeMode     string  `xml:"fee_mode" json:"fee_mode"`
	Currency    string  `xml:"currency" json:"currency"`
	Note        string  `xml:"note" json:"note"`
	Status      string  `xml:"status" json:"status"`
	Type        string  `xml:"type" json:"type"`
	Balance     float64 `xml:"balance" json:"balance"`
	Method      string  `xml:"method" json:"method"`
}

//GetHistoryRequest struct
type GetHistoryRequest struct {
	*RequestParams
	History *GetHistoryRequestParams `xml:"history" json:"history"`
}

//GetHistoryRequestParams struct
type GetHistoryRequestParams struct {
	StartDate string          `xml:"start_date,omitempty" json:"start_date"` //for specify start date. format : YYYY-mm-dd example : 2011-03-01
	EndDate   string          `xml:"end_date,omitempty" json:"end_date"`     //for specify end date. format : YYYY-mm-dd example : 2011-03-01
	Type      TransactionType `xml:"type,omitempty" json:"type"`             //for specify transaction type. (transfer|topup|redeem|exchange|receive)
	OrderBy   string          `xml:"order_by,omitempty" json:"order_by"`     //for specify order/sort by specific parameters (sorting) (date|amount|to|from|currency|bank)
	Order     string          `xml:"order,omitempty" json:"order"`           //specify order type (ASC|DESC)
	Page      uint64          `xml:"page,omitempty" json:"page"`             //for getting specific page from history transaction which has more than one page
	PageSize  uint64          `xml:"page_size,omitempty" json:"page_size"`   //for specify how much transaction per page (max 20)
}

//GetHistoryResponse struct
type GetHistoryResponse struct {
	*ResponseBody
	History *GetHistoryResponseHistoryParams `xml:"history" json:"history"`
}

//GetHistoryResponseHistoryParams struct
type GetHistoryResponseHistoryParams struct {
	Page    *GetHistoryResponsePageParams     `xml:"page" json:"page"`
	Details []*GetHistoryResponseDetailParams `xml:"detail,omitempty" json:"details,omitempty"`
}

//GetHistoryResponsePageParams struct
type GetHistoryResponsePageParams struct {
	TotalItem   uint64 `xml:"total_item" json:"total_item"`
	PageCount   uint64 `xml:"page_count" json:"page_count"`
	CurrentPage uint64 `xml:"current_page" json:"current_page"`
}

//GetHistoryResponseDetailParams struct
type GetHistoryResponseDetailParams struct {
	XMLName     xml.Name `xml:"detail" json:"-"`
	BatchNumber string   `xml:"batchnumber" json:"batchnumber"`
	Datetime    string   `xml:"datetime" json:"datetime"`
	Type        string   `xml:"type" json:"type"`
	To          string   `xml:"to" json:"to"`
	From        string   `xml:"from" json:"from"`
	Amount      float64  `xml:"amount" json:"amount"`
	Note        string   `xml:"note" json:"note"`
	Status      string   `xml:"status" json:"status"`
	Currency    string   `xml:"currency" json:"currency"`
	Fee         float64  `xml:"fee" json:"fee"`
}

//GetDetailsRequest struct
type GetDetailsRequest struct {
	*RequestParams
	Details []GetDetailsDetailParamsInterface `xml:"detail" json:"details"`
}

//GetDetailsDetailParamsInterface interface
type GetDetailsDetailParamsInterface interface {
	GetDetailType() string
}

//GetDetailsRequestDetailParamsStruct struct
type GetDetailsRequestDetailParamsStruct struct {
	XMLName xml.Name `xml:"detail" json:"-"`
	Ref     string   `xml:"ref,omitempty" json:"ref,omitempty"`   //REF parameter used to search for spesific fp_merchant_ref string that was saved by FasaPay during Transaction using SCI
	Note    string   `xml:"note,omitempty" json:"note,omitempty"` //NOTE Parameter used to search for spesific note string that was saved by FasaPay During Transaction.
}

//GetDetailType method implementation
func (f *GetDetailsRequestDetailParamsStruct) GetDetailType() string {
	return "struct"
}

//GetDetailsRequestDetailParamsString struct
type GetDetailsRequestDetailParamsString string

//GetDetailType method implementation
func (f *GetDetailsRequestDetailParamsString) GetDetailType() string {
	return "string"
}

//GetDetailsResponse struct
type GetDetailsResponse struct {
	*ResponseBody
	Details []*GetDetailsResponseDetailParams `xml:"detail,omitempty" json:"details,omitempty"`
}

//GetDetailsResponseDetailParams struct
type GetDetailsResponseDetailParams struct {
	XMLName     xml.Name `xml:"detail" json:"-"`
	Mode        string   `xml:"mode,attr" json:"mode"`
	Code        uint64   `xml:"code,attr" json:"code"`
	BatchNumber string   `xml:"batchnumber" json:"batchnumber"`
	Date        string   `xml:"date" json:"date"`
	Time        string   `xml:"time" json:"time"`
	From        string   `xml:"from" json:"from"`
	To          string   `xml:"to" json:"to"`
	Amount      float64  `xml:"amount" json:"amount"`
	Total       float64  `xml:"total" json:"total"`
	Currency    string   `xml:"currency" json:"currency"`
	Note        string   `xml:"note" json:"note"`
	Status      string   `xml:"status" json:"status"`
	Fee         float64  `xml:"fee" json:"fee"`
	Type        string   `xml:"type" json:"type"`
	Method      string   `xml:"method" json:"method"`
	FeeMode     string   `xml:"fee_mod" json:"fee_mod"`
}

//TransfersResource struct
type TransfersResource struct {
	*ResourceAbstract
}

//CreateTransfer method - allow you to transfer fund from one account to another.
//With this command you may transfer any of the available currencies that FasaPay supports.
//This function also permits you to perform multiple (bulk) transfers.
//
//Xml format of single transfer request:
//
//<fasa_request id="1234567">
//    <auth>
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <transfer id="tid">
//        <to>FP89680</to>
//        <amount>1000.0</amount>
//        <currency>idr</currency>
//        <note>standart operation</note>
//    </transfer>
//</fasa_request>
//
//Xml format of batch transfer request:
//
//<fasa_request id="1234567">
//    <auth><!-- authentication tag. required on every request -->
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <transfer id="tid-1"> <!-- transfer tag dan ididentifier -->
//        <to>FP00001</to> <!-- akun tujuan-->
//        <amount>1000.0</amount> <!-- jumlah yang ditransfer -->
//        <currency>idr</currency> <!-- kurensi yang digunakan -->
//        <note>note note</note> <!-- catatan -->
//    </transfer>
//    <transfer id="tid-2">
//        <to>FP00002</to>
//        <amount>1000.0</amount>
//        <currency>idr</currency>
//        <note>no note</note>
//    </transfer>
//    <transfer id="tid-3">
//        <to>FP00003</to>
//        <amount>1000.0</amount>
//        <currency>idr</currency>
//        <note></note>
//    </transfer>
//</fasa_request>
//
func (r *TransfersResource) CreateTransfer(transfers []*CreateTransferRequestParams, ctx context.Context, attributes *RequestParamsAttributes) (*CreateTransferResponse, *http.Response, error) {
	err := r.validateTransferParams(transfers)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.CreateTransfer error: %v", err)
	}
	baseRequestParams := r.buildRequestParams(attributes)
	requestParams := &CreateTransferRequest{baseRequestParams, transfers}
	bytesRequest, err := r.marshalRequestParams(requestParams)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.CreateTransfer error: %v", err)
	}
	rsp, err := r.tr.SendRequest(ctx, bytesRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.CreateTransfer error: %v", err)
	}
	var result CreateTransferResponse
	err = r.unmarshalResponse(rsp, &result)
	if err != nil {
		return nil, rsp, fmt.Errorf("TransfersResource.CreateTransfer error: %v", err)
	}
	if !result.IsSuccess() {
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, nil
}

//GetHistory method - allow you to receive history transaction of your FasaPay account. this command has many additional parameter to filter the response like date range, currencies, type of transaction, account target, etc.
//Request history does not need any parameter to get 10 latest transactions
//
//Basic xml format for history request:
//
//<fasa_request id="1234567">
//    <auth><!-- authentication tag. required on every request -->
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <history>
//        <start_date>2011-07-01</start_date>
//        <end_date>2011-07-09</end_date>
//        <type>transfer</type>
//        <order_by>date</order_by>
//        <order>DESC</order>
//        <page>3</page>
//        <page_size>5</page_size>
//    </history>
//</fasa_request>
//
func (r *TransfersResource) GetHistory(history *GetHistoryRequestParams, ctx context.Context, attributes *RequestParamsAttributes) (*GetHistoryResponse, *http.Response, error) {
	baseRequestParams := r.buildRequestParams(attributes)
	requestParams := &GetHistoryRequest{baseRequestParams, history}
	bytesRequest, err := r.marshalRequestParams(requestParams)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.GetHistory error: %v", err)
	}
	rsp, err := r.tr.SendRequest(ctx, bytesRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.GetHistory error: %v", err)
	}
	var result GetHistoryResponse
	err = r.unmarshalResponse(rsp, &result)
	if err != nil {
		return nil, rsp, fmt.Errorf("TransfersResource.GetHistory error: %v", err)
	}
	if !result.IsSuccess() {
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, nil
}

//GetDetails method - allow you to receive detail information of specific transaction. You can include more than one of this command in single request.
//
//Detail-Request is used to get the detailed transaction information.
//
//Detail-Request only needs BATCHNUMBER of transactions that you want to see.
//
//Detail-Request can also use this parameter to search for specific transaction:
//
//ref, REF parameter used to search for specific fp_merchant_ref string that was saved by FasaPay during Transaction using SCI
//
//note, NOTE Parameter used to search for specific note string that was saved by FasaPay During Transaction.
//
//Xml format for simple detail request:
//
//<fasa_request id="1234567">
//    <auth>
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <detail>TR2012092712345</detail>
//</fasa_request>
//
//Xml format for simple detail request:
//
//<fasa_request id="1234567">
//    <auth><!-- authentication tag. required on every request -->
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <detail>TU2012092712345</detail>
//    <detail>TR2012100265432</detail>
//    <detail>TR2012092791234</detail>
//    <detail><ref>BL12345</ref></detail>
//     <detail><note>Pembayaran</note></detail>
//</fasa_request>
//
func (r *TransfersResource) GetDetails(details []GetDetailsDetailParamsInterface, ctx context.Context, attributes *RequestParamsAttributes) (*GetDetailsResponse, *http.Response, error) {
	baseRequestParams := r.buildRequestParams(attributes)
	requestParams := &GetDetailsRequest{baseRequestParams, details}
	bytesRequest, err := r.marshalRequestParams(requestParams)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.GetDetails error: %v", err)
	}
	rsp, err := r.tr.SendRequest(ctx, bytesRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("TransfersResource.GetDetails error: %v", err)
	}
	var result GetDetailsResponse
	err = r.unmarshalResponse(rsp, &result)
	if err != nil {
		return nil, rsp, fmt.Errorf("TransfersResource.GetDetails error: %v", err)
	}
	if !result.IsSuccess() {
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, nil
}

//validateTransferParams method
func (r *TransfersResource) validateTransferParams(transfers []*CreateTransferRequestParams) error {
	var err error
	for _, transfer := range transfers {
		err = transfer.isValid()
		if err != nil {
			break
		}
	}
	return err
}

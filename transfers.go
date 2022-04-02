package fasapay

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
)

//CreateTransferRequestParams struct
type CreateTransferRequestParams struct {
	XMLName  xml.Name     `xml:"transfer"`
	Id       string       `xml:"id,attr,omitempty" json:"id"`
	To       string       `xml:"to" json:"to"`
	Amount   float64      `xml:"amount" json:"amount"`
	Currency CurrencyCode `xml:"currency" json:"currency"`
	FeeMode  string       `xml:"fee_mode,omitempty" json:"fee_mode"`
	Note     string       `xml:"note,omitempty" json:"note"`
	Ref      string       `xml:"ref,omitempty" json:"ref"`
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
	StartDate string `xml:"start_date,omitempty" json:"start_date"`
	EndDate   string `xml:"end_date,omitempty" json:"end_date"`
	Type      string `xml:"type,omitempty" json:"type"`
	OrderBy   string `xml:"order_by,omitempty" json:"order_by"`
	Order     string `xml:"order,omitempty" json:"order"`
	Page      uint64 `xml:"page,omitempty" json:"page"`
	PageSize  uint64 `xml:"page_size,omitempty" json:"page_size"`
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
	Ref     string   `xml:"ref,omitempty" json:"ref,omitempty"`
	Note    string   `xml:"note,omitempty" json:"note,omitempty"`
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

//CreateTransfer method
func (r *TransfersResource) CreateTransfer(transfers []*CreateTransferRequestParams, ctx context.Context, attributes *RequestParamsAttributes) (*CreateTransferResponse, *http.Response, error) {
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

//GetHistory method
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

//GetDetails method
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

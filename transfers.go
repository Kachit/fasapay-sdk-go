package fasapay

import "encoding/xml"

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

type CreateTransferRequest struct {
	*RequestParams
	Transfers []*CreateTransferRequestParams `xml:"transfer" json:"transfers"`
}

type CreateTransferResponse struct {
	*ResponseBody
	Transfers []*CreateTransferResponseParams `xml:"transfer" json:"transfers"`
}

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

type GetHistoryRequestParams struct {
	StartDate string `xml:"start_date,omitempty" json:"start_date"`
	EndDate   string `xml:"end_date,omitempty" json:"end_date"`
	Type      string `xml:"type,omitempty" json:"type"`
	OrderBy   string `xml:"order_by,omitempty" json:"order_by"`
	Order     string `xml:"order,omitempty" json:"order"`
	Page      uint64 `xml:"page,omitempty" json:"page"`
	PageSize  uint64 `xml:"page_size,omitempty" json:"page_size"`
}

type GetHistoryRequest struct {
	*RequestParams
	History *GetHistoryRequestParams `xml:"history" json:"history"`
}

type GetHistoryResponse struct {
	*ResponseBody
	History *GetHistoryResponseHistoryParams `xml:"history" json:"history"`
}

type GetHistoryResponseHistoryParams struct {
	Page    *GetHistoryResponsePageParams     `xml:"page" json:"page"`
	Details []*GetHistoryResponseDetailParams `xml:"detail" json:"details"`
}

type GetHistoryResponsePageParams struct {
	TotalItem   uint64 `xml:"total_item" json:"total_item"`
	PageCount   uint64 `xml:"page_count" json:"page_count"`
	CurrentPage uint64 `xml:"current_page" json:"current_page"`
}

type GetHistoryResponseDetailParams struct {
	XMLName     xml.Name `xml:"detail"`
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

type GetDetailsRequest struct {
	*RequestParams
	Details []GetDetailsDetailParamsInterface `xml:"detail" json:"details"`
}

type GetDetailsDetailParamsInterface interface {
	GetDetailType() string
}

type GetDetailsRequestDetailParamsStruct struct {
	XMLName xml.Name `xml:"detail"`
	Ref     string   `xml:"ref,omitempty" json:"ref"`
	Note    string   `xml:"note,omitempty" json:"note"`
}

func (f *GetDetailsRequestDetailParamsStruct) GetDetailType() string {
	return "struct"
}

type GetDetailsRequestDetailParamsString string

func (f *GetDetailsRequestDetailParamsString) GetDetailType() string {
	return "string"
}

type GetDetailsResponse struct {
	*ResponseBody
	Details []*GetDetailsResponseDetailParams `xml:"detail" json:"details"`
}

type GetDetailsResponseDetailParams struct {
	XMLName     xml.Name `xml:"detail"`
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

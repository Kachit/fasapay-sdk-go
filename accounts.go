package fasapay

import "encoding/xml"

type GetBalancesRequest struct {
	*RequestParams
	Balances []CurrencyCode `xml:"balance" json:"balances"`
}

type GetBalancesResponse struct {
	*ResponseBody
	Balances []*GetBalancesResponseParams `xml:"balance" json:"balances"`
}

type GetBalancesResponseParams struct {
	XMLName xml.Name `xml:"balance"`
	IDR     float64  `xml:"IDR" json:"IDR"`
	USD     float64  `xml:"USD" json:"USD"`
}

type GetAccountsRequest struct {
	*RequestParams
	Accounts []string `xml:"account" json:"accounts"`
}

type GetAccountsResponse struct {
	*ResponseBody
	Accounts []*GetAccountsResponseParams `xml:"account" json:"accounts"`
}

type GetAccountsResponseParams struct {
	XMLName  xml.Name `xml:"account"`
	FullName string   `xml:"fullname" json:"fullname"`
	Account  string   `xml:"account" json:"account"`
	Status   string   `xml:"status" json:"status"`
}

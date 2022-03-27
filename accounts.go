package fasapay

import (
	"encoding/xml"
	"net/http"
)

//GetBalancesRequest struct
type GetBalancesRequest struct {
	*RequestParams
	Balances []CurrencyCode `xml:"balance" json:"balances"`
}

//GetBalancesResponse struct
type GetBalancesResponse struct {
	*ResponseBody
	Balances []*GetBalancesResponseParams `xml:"balance,omitempty" json:"balances,omitempty"`
}

//GetBalancesResponseParams struct
type GetBalancesResponseParams struct {
	XMLName xml.Name `xml:"balance" json:"-"`
	IDR     float64  `xml:"IDR" json:"IDR"`
	USD     float64  `xml:"USD" json:"USD"`
}

//GetAccountsRequest struct
type GetAccountsRequest struct {
	*RequestParams
	Accounts []string `xml:"account" json:"accounts"`
}

//GetAccountsResponse struct
type GetAccountsResponse struct {
	*ResponseBody
	Accounts []*GetAccountsResponseParams `xml:"account,omitempty" json:"accounts,omitempty"`
}

//GetAccountsResponseParams struct
type GetAccountsResponseParams struct {
	XMLName  xml.Name `xml:"account" json:"-"`
	FullName string   `xml:"fullname" json:"fullname"`
	Account  string   `xml:"account" json:"account"`
	Status   string   `xml:"status" json:"status"`
}

type AccountsResource struct {
	*ResourceAbstract
}

func (r *AccountsResource) GetBalances(currencies []CurrencyCode, custom *CustomRequestParams) (*GetBalancesResponse, *http.Response, error) {
	return nil, nil, nil
}

func (r *AccountsResource) GetAccounts(accounts []string, custom *CustomRequestParams) (*GetAccountsResponse, *http.Response, error) {
	return nil, nil, nil
}

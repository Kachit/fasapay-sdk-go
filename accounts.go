package fasapay

import (
	"context"
	"encoding/xml"
	"fmt"
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
	Balances *GetBalancesResponseParams `xml:"balance,omitempty" json:"balances,omitempty"`
}

//GetBalancesResponseParams struct
type GetBalancesResponseParams struct {
	IDR float64 `xml:"IDR" json:"IDR"`
	USD float64 `xml:"USD" json:"USD"`
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

func (r *AccountsResource) GetBalances(currencies []CurrencyCode, ctx context.Context, attributes *RequestParamsAttributes) (*GetBalancesResponse, *http.Response, error) {
	baseRequestParams := r.buildRequestParams(attributes)
	requestParams := &GetBalancesRequest{baseRequestParams, currencies}
	bytesRequest, err := r.marshalRequestParams(requestParams)
	if err != nil {
		return nil, nil, fmt.Errorf("AccountsResource.GetBalances error: %v", err)
	}
	rsp, err := r.tr.SendRequest(ctx, bytesRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("AccountsResource.GetBalances error: %v", err)
	}
	var result GetBalancesResponse
	err = r.unmarshalResponse(rsp, &result)
	if err != nil {
		return nil, rsp, fmt.Errorf("AccountsResource.GetBalances error: %v", err)
	}
	return &result, rsp, nil
}

func (r *AccountsResource) GetAccounts(accounts []string, ctx context.Context, attributes *RequestParamsAttributes) (*GetAccountsResponse, *http.Response, error) {
	baseRequestParams := r.buildRequestParams(attributes)
	requestParams := &GetAccountsRequest{baseRequestParams, accounts}
	bytesRequest, err := r.marshalRequestParams(requestParams)
	if err != nil {
		return nil, nil, fmt.Errorf("AccountsResource.GetAccounts error: %v", err)
	}
	rsp, err := r.tr.SendRequest(ctx, bytesRequest)
	if err != nil {
		return nil, nil, fmt.Errorf("AccountsResource.GetAccounts error: %v", err)
	}
	var result GetAccountsResponse
	err = r.unmarshalResponse(rsp, &result)
	if err != nil {
		return nil, rsp, fmt.Errorf("AccountsResource.GetAccounts error: %v", err)
	}
	return &result, rsp, nil
}

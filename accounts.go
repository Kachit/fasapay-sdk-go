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

//AccountsResource struct
type AccountsResource struct {
	*ResourceAbstract
}

//GetBalances method - allow you to check your FasaPay account balance.
//Balance request is used to get the amount of balance in your account.
//Balance request only needs currency code of currency that you want to see.  (IDR, USD)
//xml format for single balance request:
//<fasa_request id="1234567">
//    <auth>
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <balance>IDR</balance>
//</fasa_request>
//xml format for batch balances request:
//<fasa_request id="1234567">
//    <auth><!-- authentication tag. required on every request -->
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <balance>IDR</balance>
//    <balance>USD</balance>
//</fasa_request>
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
	if !result.IsSuccess() {
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, nil
}

//GetAccounts method - allow you to check specific FasaPay account, to indicate is it registered or not.
//Account request is used to get the information of certain FasaPay user by their account number.
//Account request only needs account number of FasaPay account that you want to see.
//basic xml format for single account request:
//<fasa_request id="1234567">
//    <auth>
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <account>FP00001</account>
//</fasa_request>
//basic xml format for batch accounts request:
//<fasa_request id="1234567">
//    <auth><!-- authentication tag. required on every request -->
//        <api_key>11123548cd3a5e5613325132112becf</api_key>
//        <token>e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b</token>
//    </auth>
//    <account>FP00001</account>
//    <account>FP00002</account>
//</fasa_request>
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
	if !result.IsSuccess() {
		return &result, rsp, fmt.Errorf(result.GetError())
	}
	return &result, rsp, nil
}

package fasapay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//RequestBuilder handler
type RequestBuilder struct {
	cfg *Config
}

//BuildAuthParams method
func (rb *RequestBuilder) buildAuthParams(dt time.Time) *RequestAuthParams {
	params := &RequestAuthParams{
		ApiKey: rb.cfg.ApiKey,
		Token:  generateAuthToken(rb.cfg.ApiKey, rb.cfg.ApiSecretWord, dt),
	}
	return params
}

//BuildParams method
func (rb *RequestBuilder) buildParams(id string, dt time.Time) *RequestParams {
	params := &RequestParams{Id: id, Auth: rb.buildAuthParams(dt)}
	return params
}

//BuildQueryParams method
func (rb *RequestBuilder) buildQueryParams(query map[string]interface{}) string {
	q := url.Values{}
	if query != nil {
		for k, v := range query {
			q.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return q.Encode()
}

//NewHttpTransport create new http transport
func NewHttpTransport(config *Config, h *http.Client) *Transport {
	rb := &RequestBuilder{cfg: config}
	return &Transport{http: h, rb: rb}
}

//Transport wrapper
type Transport struct {
	http *http.Client
	rb   *RequestBuilder
}

//UnmarshalResponse func
func unmarshalResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Response.Unmarshal read body: %v", err)
	}
	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return xml.Unmarshal(bodyBytes, &v)
}

//CustomRequestParams struct
type CustomRequestParams struct {
	Id       string    `json:"id"`
	DateTime time.Time `json:"date_time"`
}

//RequestParams struct
type RequestParams struct {
	XMLName xml.Name           `xml:"fasa_request"`
	Id      string             `xml:"id,attr"`
	Auth    *RequestAuthParams `xml:"auth" json:"auth"`
}

//RequestAuthParams struct
type RequestAuthParams struct {
	XMLName xml.Name `xml:"auth" json:"-"`
	ApiKey  string   `xml:"api_key" json:"api_key"`
	Token   string   `xml:"token" json:"token"`
}

//ResponseBody struct
type ResponseBody struct {
	XMLName  xml.Name            `xml:"fasa_response" json:"-"`
	Id       string              `xml:"id,attr" json:"id"`
	DateTime string              `xml:"date_time,attr" json:"date_time"`
	Errors   *ResponseBodyErrors `xml:"errors,omitempty" json:"errors,omitempty"`
}

//IsSuccess method
func (r *ResponseBody) IsSuccess() bool {
	return r.Errors == nil
}

//ResponseBodyErrors struct
type ResponseBodyErrors struct {
	XMLName xml.Name                   `xml:"errors" json:"-"`
	Id      string                     `xml:"id,attr,omitempty" json:"id,omitempty"`
	Mode    string                     `xml:"mode,attr" json:"mode"`
	Code    uint64                     `xml:"code,attr" json:"code"`
	Data    []*ResponseBodyErrorParams `xml:"data" json:"data"`
}

//ResponseBodyErrorParams struct
type ResponseBodyErrorParams struct {
	XMLName   xml.Name `xml:"data" json:"-"`
	Code      uint64   `xml:"code,omitempty" json:"code,omitempty"`
	Attribute string   `xml:"attribute,omitempty" json:"attribute,omitempty"`
	Message   string   `xml:"message,omitempty" json:"message,omitempty"`
	Detail    string   `xml:"detail,omitempty" json:"detail,omitempty"`
}

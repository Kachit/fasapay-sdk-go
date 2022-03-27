package fasapay

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

//RequestBuilder handler
type RequestBuilder struct {
	cfg *Config
}

//BuildUri method
func (rb *RequestBuilder) buildUri() (uri *url.URL, err error) {
	u, err := url.Parse(rb.cfg.Uri)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder.buildUri parse: %v", err)
	}
	return u, err
}

//BuildBody method
func (rb *RequestBuilder) buildBody(body []byte) (io.Reader, error) {
	return bytes.NewBuffer(body), nil
}

//BuildRequest method
func (rb *RequestBuilder) buildRequest(ctx context.Context, body []byte) (req *http.Request, err error) {
	//build body
	bodyReader, err := rb.buildBody(body)
	//build uri
	uri, err := rb.buildUri()
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder.buildRequest build uri: %v", err)
	}
	//build request
	req, err = http.NewRequestWithContext(ctx, "POST", uri.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("RequestBuilder.buildRequest new request error: %v", err)
	}
	return req, nil
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

//SendRequest Send request method
func (tr *Transport) SendRequest(ctx context.Context, body []byte) (resp *http.Response, err error) {
	req, err := tr.rb.buildRequest(ctx, body)
	if err != nil {
		return nil, fmt.Errorf("transport.SendRequest: %v", err)
	}
	return tr.http.Do(req)
}

//RequestParamsAttributes struct
type RequestParamsAttributes struct {
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

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

//ResourceAbstract base resource
type ResourceAbstract struct {
	tr  *Transport
	cfg *Config
}

//BuildAuthParams method
func (rb *ResourceAbstract) buildAuthRequestParams(dt time.Time) *RequestAuthParams {
	params := &RequestAuthParams{
		ApiKey: rb.cfg.ApiKey,
		Token:  generateAuthToken(rb.cfg.ApiKey, rb.cfg.ApiSecretWord, dt),
	}
	return params
}

//BuildParams method
func (rb *ResourceAbstract) buildRequestParams(attributes *RequestParamsAttributes) *RequestParams {
	if attributes == nil {
		dt := time.Now().UTC()
		attributes = &RequestParamsAttributes{Id: fmt.Sprint(dt.Unix()), DateTime: dt}
	}
	params := &RequestParams{Id: attributes.Id, Auth: rb.buildAuthRequestParams(attributes.DateTime)}
	return params
}

//UnmarshalResponse func
func (rb *ResourceAbstract) marshalRequestParams(request interface{}) ([]byte, error) {
	bts, err := xml.Marshal(request)
	if err != nil {
		return nil, err
	}
	req := "req=" + url.QueryEscape(string(bts))
	return []byte(req), nil
}

//UnmarshalResponse func
func (rb *ResourceAbstract) unmarshalResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Response.Unmarshal read body: %v", err)
	}
	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return xml.Unmarshal(bodyBytes, &v)
}

//NewResourceAbstract Create new resource abstract
func NewResourceAbstract(transport *Transport, config *Config) *ResourceAbstract {
	return &ResourceAbstract{tr: transport, cfg: config}
}

package fasapay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//ResourceAbstract base resource
type ResourceAbstract struct {
	tr  *Transport
	cfg *Config
}

//BuildAuthParams method
func (ra *ResourceAbstract) buildAuthRequestParams(dt time.Time) *RequestAuthParams {
	params := &RequestAuthParams{
		ApiKey: ra.cfg.ApiKey,
		Token:  generateAuthToken(ra.cfg.ApiKey, ra.cfg.ApiSecretWord, dt),
	}
	return params
}

//BuildParams method
func (ra *ResourceAbstract) buildRequestParams(attributes *RequestParamsAttributes) *RequestParams {
	if attributes == nil {
		dt := time.Now().UTC()
		attributes = &RequestParamsAttributes{Id: fmt.Sprint(dt.Unix()), DateTime: dt}
	}
	params := &RequestParams{Id: attributes.Id, Auth: ra.buildAuthRequestParams(attributes.DateTime)}
	return params
}

//UnmarshalResponse func
func (ra *ResourceAbstract) marshalRequestParams(request interface{}) ([]byte, error) {
	bts, err := xml.Marshal(request)
	if err != nil {
		return nil, err
	}
	req := []byte("req=")
	req = append(req, bts...)
	return req, nil
}

//UnmarshalResponse func
func (ra *ResourceAbstract) unmarshalResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ResourceAbstract.unmarshalResponse read body: %v", err)
	}
	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return xml.Unmarshal(bodyBytes, &v)
}

//NewResourceAbstract Create new resource abstract
func NewResourceAbstract(transport *Transport, config *Config) *ResourceAbstract {
	return &ResourceAbstract{tr: transport, cfg: config}
}

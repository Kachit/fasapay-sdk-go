package fasapay

import "encoding/xml"

type RequestParams struct {
	XMLName xml.Name           `xml:"fasa_request"`
	Id      string             `xml:"id,attr"`
	Auth    *RequestAuthParams `xml:"auth" json:"auth"`
}

type RequestAuthParams struct {
	XMLName xml.Name `xml:"auth"`
	ApiKey  string   `xml:"api_key" json:"api_key"`
	Token   string   `xml:"token" json:"token"`
}

type ResponseBody struct {
	XMLName  xml.Name            `xml:"fasa_response"`
	Id       string              `xml:"id,attr" json:"id"`
	DateTime string              `xml:"date_time,attr" json:"date_time"`
	Errors   *ResponseBodyErrors `xml:"errors" json:"errors"`
}

type ResponseBodyErrors struct {
	XMLName xml.Name                   `xml:"errors"`
	Id      string                     `xml:"id,attr" json:"id"`
	Mode    string                     `xml:"mode,attr" json:"mode"`
	Code    uint64                     `xml:"code,attr" json:"code"`
	Data    []*ResponseBodyErrorParams `xml:"data" json:"data"`
}

type ResponseBodyErrorParams struct {
	XMLName   xml.Name `xml:"data"`
	Code      uint64   `xml:"code" json:"code"`
	Attribute string   `xml:"attribute" json:"attribute"`
	Message   string   `xml:"message" json:"message"`
	Detail    string   `xml:"detail" json:"detail"`
}

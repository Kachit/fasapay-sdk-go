package fasapay

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuildStubConfig() *Config {
	return &Config{
		ApiUri:        SandboxAPIUrl,
		ApiKey:        "ApiKey",
		ApiSecretWord: "ApiSecretWord",
	}
}

func BuildStubRequest() *RequestParams {
	auth := &RequestAuthParams{ApiKey: "foo", Token: "bar"}
	return &RequestParams{Auth: auth, Id: "123456"}
}

func LoadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func BuildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func BuildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := LoadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}

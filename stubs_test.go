package fasapay

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const TestableApiKey = "11123548cd3a5e5613325132112becf"
const TestableApiSecretWord = "kata rahasia"
const TestableApiAuthToken = "e910361e42dafdfd100b19701c2ef403858cab640fd699afc67b78c7603ddb1b"

//BuildStubConfig method
func BuildStubConfig() *Config {
	return &Config{
		Uri:           SandboxAPIUrl,
		ApiKey:        TestableApiKey,
		ApiSecretWord: TestableApiSecretWord,
	}
}

//BuildStubHttpTransport method
func BuildStubHttpTransport() *Transport {
	return NewHttpTransport(BuildStubConfig(), &http.Client{})
}

//BuildStubDateTime method
func BuildStubDateTime() time.Time {
	return time.Date(2011, time.Month(7), 20, 15, 30, 0, 0, time.UTC)
}

//BuildStubRequest method
func BuildStubRequest() *RequestParams {
	auth := &RequestAuthParams{ApiKey: TestableApiKey, Token: TestableApiAuthToken}
	return &RequestParams{Auth: auth, Id: "1234567"}
}

//LoadStubResponseData method
func LoadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

//BuildStubResponseFromString method
func BuildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

//BuildStubResponseFromFile method
func BuildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := LoadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}

package fasapay

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri()
	assert.NotEmpty(t, uri)
	assert.Nil(t, err)
	assert.Equal(t, SandboxAPIUrl, uri.String())
}

func Test_HTTP_RequestBuilder_BuildHeaders(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	headers := builder.buildHeaders()
	assert.NotEmpty(t, headers)
	assert.Equal(t, "application/x-www-form-urlencoded", headers.Get("Content-Type"))
}

//func Test_HTTP_RequestBuilder_BuildBody(t *testing.T) {
//	cfg := BuildStubConfig()
//	builder := RequestBuilder{cfg: cfg}
//
//	data := make(map[string]interface{})
//	data["foo"] = "bar"
//	data["bar"] = "baz"
//
//	body, _ := builder.buildBody(data)
//	assert.NotEmpty(t, body)
//}

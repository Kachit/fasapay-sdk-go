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

func Test_HTTP_RequestBuilder_BuildBody(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := "foo"
	body, _ := builder.buildBody([]byte(data))
	assert.NotEmpty(t, body)
}

func Test_HTTP_ResponseBody_IsSuccess(t *testing.T) {
	rsp := &ResponseBody{}
	assert.True(t, rsp.IsSuccess())
	rsp.Errors = &ResponseBodyErrors{}
	assert.False(t, rsp.IsSuccess())
}

func Test_HTTP_ResponseBody_GetError(t *testing.T) {
	rsp := &ResponseBody{Errors: &ResponseBodyErrors{}}
	assert.Equal(t, ErrorMessageUnexpectedError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeNotValidXmlRequest
	assert.Equal(t, ErrorMessageNotValidXmlRequest, rsp.GetError())
	rsp.Errors.Code = ErrorCodeUnauthorized
	assert.Equal(t, ErrorMessageUnauthorized, rsp.GetError())
	rsp.Errors.Code = ErrorCodeNotAcceptableTransfer
	assert.Equal(t, ErrorMessageNotAcceptableTransfer, rsp.GetError())
	rsp.Errors.Code = ErrorCodeDetailRequestError
	assert.Equal(t, ErrorMessageDetailRequestError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeHistoryRequestError
	assert.Equal(t, ErrorMessageHistoryRequestError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeBalanceRequestError
	assert.Equal(t, ErrorMessageBalanceRequestError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeAccountRequestError
	assert.Equal(t, ErrorMessageAccountRequestError, rsp.GetError())
}

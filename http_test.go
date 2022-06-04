package fasapay

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type HttpRequestBuilderTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *RequestBuilder
}

func (suite *HttpRequestBuilderTestSuite) SetupTest() {
	suite.cfg = BuildStubConfig()
	suite.ctx = context.Background()
	suite.testable = &RequestBuilder{cfg: suite.cfg}
}

func (suite *HttpRequestBuilderTestSuite) TestBuildUriWithoutQueryParams() {
	uri, err := suite.testable.buildUri()
	assert.NotEmpty(suite.T(), uri)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), SandboxAPIUrl, uri.String())
}

func (suite *HttpRequestBuilderTestSuite) TestBuildHeaders() {
	headers := suite.testable.buildHeaders()
	assert.NotEmpty(suite.T(), headers)
	assert.Equal(suite.T(), "application/x-www-form-urlencoded", headers.Get("Content-Type"))
}

func (suite *HttpRequestBuilderTestSuite) TestBuildBody() {
	data := "foo"
	body, err := suite.testable.buildBody([]byte(data))
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), body)
}

func TestHttpRequestBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(HttpRequestBuilderTestSuite))
}

type HttpResponseBodyTestSuite struct {
	suite.Suite
}

func (suite *HttpResponseBodyTestSuite) TestIsSuccess() {
	rsp := &ResponseBody{}
	assert.True(suite.T(), rsp.IsSuccess())
	rsp.Errors = &ResponseBodyErrors{}
	assert.False(suite.T(), rsp.IsSuccess())
}

func (suite *HttpResponseBodyTestSuite) TestGetError() {
	rsp := &ResponseBody{Errors: &ResponseBodyErrors{}}
	assert.Equal(suite.T(), ErrorMessageUnexpectedError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeNotValidXmlRequest
	assert.Equal(suite.T(), ErrorMessageNotValidXmlRequest, rsp.GetError())
	rsp.Errors.Code = ErrorCodeUnauthorized
	assert.Equal(suite.T(), ErrorMessageUnauthorized, rsp.GetError())
	rsp.Errors.Code = ErrorCodeNotAcceptableTransfer
	assert.Equal(suite.T(), ErrorMessageNotAcceptableTransfer, rsp.GetError())
	rsp.Errors.Code = ErrorCodeDetailRequestError
	assert.Equal(suite.T(), ErrorMessageDetailRequestError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeHistoryRequestError
	assert.Equal(suite.T(), ErrorMessageHistoryRequestError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeBalanceRequestError
	assert.Equal(suite.T(), ErrorMessageBalanceRequestError, rsp.GetError())
	rsp.Errors.Code = ErrorCodeAccountRequestError
	assert.Equal(suite.T(), ErrorMessageAccountRequestError, rsp.GetError())
}

func TestHttpResponseBodyTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseBodyTestSuite))
}

type HttpTransportTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *Transport
}

func (suite *HttpTransportTestSuite) SetupTest() {
	suite.cfg = BuildStubConfig()
	suite.ctx = context.Background()
	suite.testable = BuildStubHttpTransport()
	httpmock.Activate()
}

func (suite *HttpTransportTestSuite) TestSendRequest() {
	body, _ := LoadStubResponseData("stubs/accounts/balances/success.xml")

	httpmock.RegisterResponder(http.MethodPost, suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, err := suite.testable.SendRequest(suite.ctx, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)

	defer resp.Body.Close()
	bodyRsp, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(suite.T(), body, bodyRsp)
}

func (suite *HttpTransportTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

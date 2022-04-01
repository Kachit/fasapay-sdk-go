package fasapay

const (
	ErrorCodeNotValidXmlRequest    uint64 = 40000
	ErrorCodeUnauthorized          uint64 = 40100
	ErrorCodeNotAcceptableTransfer uint64 = 40600
	ErrorCodeDetailRequestError    uint64 = 40700
	ErrorCodeHistoryRequestError   uint64 = 40800
	ErrorCodeBalanceRequestError   uint64 = 40900
	ErrorCodeAccountRequestError   uint64 = 41000
)

const (
	ErrorMessageNotValidXmlRequest    string = "NOT VALID XML REQUEST"
	ErrorMessageUnauthorized          string = "UNAUTHORIZED"
	ErrorMessageNotAcceptableTransfer string = "NOT ACCEPTABLE TRANSFER"
	ErrorMessageDetailRequestError    string = "DETAIL REQUEST ERROR"
	ErrorMessageHistoryRequestError   string = "HISTORY REQUEST ERROR"
	ErrorMessageBalanceRequestError   string = "BALANCE REQUEST ERROR"
	ErrorMessageAccountRequestError   string = "ACCOUNT REQUEST ERROR"
	ErrorMessageUnexpectedError       string = "UNEXPECTED ERROR"
)

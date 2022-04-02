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
	ErrorMessageNotValidXmlRequest    string = "NOT VALID XML REQUEST"   //The sended XML are not valid, broken or has wrong format
	ErrorMessageUnauthorized          string = "UNAUTHORIZED"            //Authorisation failed.
	ErrorMessageNotAcceptableTransfer string = "NOT ACCEPTABLE TRANSFER" //There is an error in the transfer operation
	ErrorMessageDetailRequestError    string = "DETAIL REQUEST ERROR"    //There is an error in the detail operation
	ErrorMessageHistoryRequestError   string = "HISTORY REQUEST ERROR"   //There is an error in the history operation
	ErrorMessageBalanceRequestError   string = "BALANCE REQUEST ERROR"   //There is an error in the balance operation
	ErrorMessageAccountRequestError   string = "ACCOUNT REQUEST ERROR"   //There is an error in the account operation
	ErrorMessageUnexpectedError       string = "UNEXPECTED ERROR"        //Unexpected error
)

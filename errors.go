package fasapay

const (
	//ErrorCodeNotValidXmlRequest The sent XML are not valid, broken or has wrong format
	ErrorCodeNotValidXmlRequest uint64 = 40000
	//ErrorCodeUnauthorized Authorisation failed.
	ErrorCodeUnauthorized uint64 = 40100
	//ErrorCodeNotAcceptableTransfer //There is an error in the transfer operation
	ErrorCodeNotAcceptableTransfer uint64 = 40600
	//ErrorCodeDetailRequestError //There is an error in the detail operation
	ErrorCodeDetailRequestError uint64 = 40700
	//ErrorCodeHistoryRequestError //There is an error in the history operation
	ErrorCodeHistoryRequestError uint64 = 40800
	//ErrorCodeBalanceRequestError //There is an error in the balance operation
	ErrorCodeBalanceRequestError uint64 = 40900
	//ErrorMessageAccountRequestError //There is an error in the account operation
	ErrorCodeAccountRequestError uint64 = 41000
)

const (
	//ErrorMessageNotValidXmlRequest The sent XML are not valid, broken or has wrong format
	ErrorMessageNotValidXmlRequest string = "NOT VALID XML REQUEST"
	//ErrorMessageUnauthorized Authorisation failed.
	ErrorMessageUnauthorized string = "UNAUTHORIZED"
	//ErrorMessageNotAcceptableTransfer  There is an error in the transfer operation
	ErrorMessageNotAcceptableTransfer string = "NOT ACCEPTABLE TRANSFER"
	//ErrorMessageDetailRequestError There is an error in the detail operation
	ErrorMessageDetailRequestError string = "DETAIL REQUEST ERROR"
	//ErrorMessageHistoryRequestError There is an error in the history operation
	ErrorMessageHistoryRequestError string = "HISTORY REQUEST ERROR"
	//ErrorMessageBalanceRequestError There is an error in the balance operation
	ErrorMessageBalanceRequestError string = "BALANCE REQUEST ERROR"
	//ErrorMessageAccountRequestError There is an error in the account operation
	ErrorMessageAccountRequestError string = "ACCOUNT REQUEST ERROR"
	//ErrorMessageUnexpectedError Unexpected error
	ErrorMessageUnexpectedError string = "UNEXPECTED ERROR"
)

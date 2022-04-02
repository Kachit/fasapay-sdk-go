package fasapay

//CurrencyCode type
type CurrencyCode string

//CurrencyCodeUSD const
const CurrencyCodeUSD CurrencyCode = "USD"

//CurrencyCodeUSD const
const CurrencyCodeIDR CurrencyCode = "IDR"

//TransactionFeeMode type
type TransactionFeeMode string

//TransactionFeeModeFiR const
const TransactionFeeModeFiR TransactionFeeMode = "FiR"

//TransactionFeeModeFiS const
const TransactionFeeModeFiS TransactionFeeMode = "FiS"

//TransactionStatus type
type TransactionStatus string

//TransactionStatusFinish const
const TransactionStatusFinish TransactionStatus = "FINISH"

//TransactionType type
type TransactionType string

//TransactionTypeTransfer const
const TransactionTypeTransfer = "transfer"

//TransactionTypeTopUp const
const TransactionTypeTopUp = "topup"

//TransactionTypeRedeem const
const TransactionTypeRedeem = "redeem"

//TransactionTypeExchange const
const TransactionTypeExchange = "exchange"

//TransactionTypeReceive const
const TransactionTypeReceive = "receive"

package paytm

import "os"

var (
	PaytmMerchantKey     = os.Getenv("PAYTM_MERCHANT_KEY")
	MID                  = os.Getenv("PAYTM_MID")
	INDUSTRY_TYPE_ID     = `Retail`
	CHANNEL_ID           = `WAP`
	WEBSITE              = os.Getenv("PAYTM_WEBSITE")
	CALLBACK_URL         = os.Getenv("PAYTM_PAYMENT_CALLBACK")
	TransactionStatusAPI = `https://securegw-stage.paytm.in/merchant-status/getTxnStatus`
)

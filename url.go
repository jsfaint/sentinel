package main

const (
	apiLoginURL        = "https://account.onethingpcs.com/user/login?appversion=1.4.8"
	apiIncomeURL       = "https://account.onethingpcs.com/wkb/income-history"
	apiOutcomeURL      = "https://account.onethingpcs.com/wkb/outcome-history"
	apiCheckAccountURL = "https://account.onethingpcs.com/user/check?appversion=1.4.8"
	apiCoinInfoURL     = "http://account.onethingpcs.com/info/query"
	apiUserURL         = "https://account.onethingpcs.com/activiate/userinfo"
	apiAccountInfoURL  = "https://account.onethingpcs.com/wkb/account-info"
)

var (
	headers = map[string]string{
		"user-agent":    "Mozilla/5.0",
		"cache-control": "no-cache",
	}
)

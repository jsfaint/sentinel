package main

const (
	apiAccountURL  = "https://account.onethingpcs.com"
	apiControlURL  = "https://control.onethingpcs.com"
	apiRemoteDlURL = "https://control.remotedl.onethingpcs.com"
	apiLoginURL    = "https://account.onethingpcs.com/user/login?appversion=1.4.8"
	apiIncomeURL   = "https://account.onethingpcs.com/wkb/income-history"
	apiOutcomeURL  = "https://account.onethingpcs.com/wkb/outcome-history"
)

var (
	headers = map[string]string{"user-agent": "Mozilla/5.0"}
)

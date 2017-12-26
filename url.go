package main

const (
	apiAccountUrl  = "https://account.onethingpcs.com"
	apiControlUrl  = "https://control.onethingpcs.com"
	apiRemoteDlUrl = "https://control.remotedl.onethingpcs.com"
	apiLoginUrl    = "https://account.onethingpcs.com/user/login?appversion=1.4.8"
	apiIncomeUrl   = "https://account.onethingpcs.com/wkb/income-history"
	apiOutcomeUrl  = "https://account.onethingpcs.com/wkb/outcome-history"
)

var (
	headers = map[string]string{"user-agent": "Mozilla/5.0"}
)

package main

const (
	apiAccountURL = "https://account.onethingpcs.com"
	apiControlURL = "https://control.onethingpcs.com"
	apiRemoteURL  = "https://control.remotedl.onethingpcs.com"

	//Account
	apiLoginURL        = apiAccountURL + "/user/login"
	apiIncomeURL       = apiAccountURL + "/wkb/income-history"
	apiOutcomeURL      = apiAccountURL + "/wkb/outcome-history"
	apiCheckAccountURL = apiAccountURL + "/user/check"
	apiCoinInfoURL     = apiAccountURL + "/info/query"
	apiUserURL         = apiAccountURL + "/activiate/userinfo"
	apiAccountInfoURL  = apiAccountURL + "/wkb/account-info"
	apiActivateInfoURL = apiAccountURL + "/activate/userinfo"

	//Control
	apiListPeerURL = apiControlURL + "/ListPeer"
	apiDiskInfoURL = apiControlURL + "/getUSBInfo"

	//Remote
	apiDownloadInfoURL = apiRemoteURL + "list"
)

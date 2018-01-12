package sentinel

const (
	apiAccountURL = "http://account.onethingpcs.com"
	apiControlURL = "http://control.onethingpcs.com"
	apiRemoteURL  = "http://control.remotedl.onethingpcs.com"

	//Account
	apiAccountInfoURL  = apiAccountURL + "/wkb/account-info"
	apiActivateInfoURL = apiAccountURL + "/activate/userinfo"
	apiCheckAccountURL = apiAccountURL + "/user/check"
	apiCoinInfoURL     = apiAccountURL + "/info/query"
	apiIncomeURL       = apiAccountURL + "/wkb/income-history"
	apiLoginURL        = apiAccountURL + "/user/login"
	apiOutcomeURL      = apiAccountURL + "/wkb/outcome-history"
	apiUserURL         = apiAccountURL + "/activiate/userinfo"
	apiWKBDrawURL      = apiAccountURL + "/wkb/draw"
	apiSessionURL      = apiAccountURL + "/user/check-session"

	//Control
	apiUSBInfoURL  = apiControlURL + "/getUSBInfo"
	apiListPeerURL = apiControlURL + "/listPeer"

	//Remote
	apiDownloadInfoURL = apiRemoteURL + "list"
)

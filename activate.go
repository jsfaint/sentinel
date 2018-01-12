package sentinel

import (
	"net/http"

	"github.com/imroc/req"
)

/*
{
    "iRet": 0,
    "sMsg": "ok",
    "data": {
        "activate_days": 16,
        "yes_wkb": 0,
        "is_done": 0,
        "type": 6
    }
}
*/

type respActivateInfo struct {
	respHead
	Data activateInfo `json:"data"`
}

type activateInfo struct {
	ActivateDays int     `json:"activate_days"`
	YesWKB       float64 `json:"yes_wkb"`
	IsDone       int     `json:"is_done"`
	Type         int     `json:"type"`
}

func (user *UserReq) getActivate() (err error) {
	r := user.r

	device := user.Peers.Devices[0]

	sign := getSign(false, map[string]string{
		"sn": device.DeviceSn,
	}, user.sessionID)

	query := req.QueryParam{
		"appversion": wkAppVersion,
	}

	body := req.Param{
		"sn":   device.DeviceSn,
		"sign": sign,
	}

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:  "origin",
			Value: "1",
		},
		&http.Cookie{
			Name:  "peerid",
			Value: device.Peerid,
		},
	}

	resp, err := r.Post(apiActivateInfoURL, headers, query, body, cookies)
	if err != nil {
		return err
	}

	var v respActivateInfo
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return v
	}

	user.ActivateInfo = &v.Data

	return
}

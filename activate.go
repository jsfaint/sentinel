package main

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
	ActivateDays int `json:"activate_days"`
	YesWKB       int `json:"yes_wkb"`
	IsDone       int `json:"is_done"`
	Type         int `json:"type"`
}

func (user *userReq) getActivate() (err error) {
	r := user.r

	query := req.QueryParam{
		"appversion": appVersion,
	}

	//FIXME
	body := req.Param{
		"sign": "",
		"sn":   "",
	}

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:  "origin",
			Value: "1",
		},
		&http.Cookie{
			Name:  "peerid",
			Value: "1", //FIXME: user.peerID
		},
	}

	resp, err := r.Post(apiAccountInfoURL, headers, query, body, cookies)
	if err != nil {
		return err
	}

	var v respActivateInfo
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return errActivateInfo
	}

	user.activateInfo = &v.Data

	return
}

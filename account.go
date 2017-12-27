package main

import (
	"net/http"

	"github.com/imroc/req"
)

/*
{
  "iRet": 0,
  "sMsg": "",
  "data": {
    "fronzen_time": 0,
    "bind_time": 1513127124,
    "balance": "",
    "isBindAddr": 1,
    "addr": "",
    "gasInfo": {
      "fastGas": "",
      "normalGas": "",
      "isFree": 0
    }
  }
}
*/
type accountInfoResp struct {
	respHead
	Data accountInfo `json:"data"`
}

type accountInfo struct {
	Addr        string  `json:"addr"`
	Balance     string  `json:"balance"`
	BindTime    uint64  `json:"bind_time"`
	FronzenTime uint64  `json:"fronzen_time"`
	GasInfo     gasInfo `json:"gasInfo"`
	IsBindAddr  uint8   `json:"isBindAddr"`
}

type gasInfo struct {
	FastGas   float64 `json:"fast_gas"`
	NormalGas float64 `json:"normal_gas"`
	IsFree    bool    `json:"is_free"`
}

/*
getAccountInfo
*/
func (user *userReq) getAccountInfo() (err error) {
	r := user.r

	//POST query parameter
	param := req.Param{
		"appversion": appVersion,
	}

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:  "sessionid",
			Value: user.sessionid,
		},
		&http.Cookie{
			Name:  "userid",
			Value: user.userid,
		},
		&http.Cookie{
			Name:  "origin",
			Value: "1",
		},
	}

	resp, err := r.Post(apiAccountInfoURL, headers, param, cookies)
	if err != nil {
		return err
	}

	var v accountInfoResp
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return errAccountInfo
	}

	user.accountInfo = &v.Data

	return
}

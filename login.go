package main

import (
	"fmt"
	"github.com/imroc/req"
)

/*
{
  "iRet": 0,
  "sMsg": "ok",
  "data": {
    "userid": "",
    "phone": "",
    "account_type": "",
    "bind_pwd": 1,
    "nickname": ""
  }
}
*/

type loginResp struct {
	respHead
	Data userInfo `json:"data"`
}

type userInfo struct {
	Userid      string `json:"userid"`
	Phone       string `json:"phone"`
	AccountType string `json:"account_type"`
	BindPwd     uint   `json:"bind_pwd"`
	NickName    string `json:"nickname"`
}

/*
login
*/
func login(r *req.Req, phone, pwd, dev, imei, sign string) (sessionid, userid string, err error) {
	//POST query parameter
	param := req.Param{
		"deviceid":     dev,
		"imeiid":       imei,
		"phone":        phone,
		"pwd":          pwd,
		"account_type": "4",
		"sign":         sign,
	}

	resp, err := r.Post(apiLoginURL, headers, param)
	if err != nil {
		return "", "", err
	}

	for _, v := range resp.Response().Cookies() {
		switch v.Name {
		case "sessionid":
			sessionid = v.Value
		case "userid":
			userid = v.Value
		}
	}

	var v loginResp
	if err := resp.ToJSON(&v); err != nil {
		return "", "", err
	}

	if !v.success() {
		return "", "", ERR_LOGIN
	}

	fmt.Println(v.Data, "\n")

	return
}

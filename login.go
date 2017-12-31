package main

import "github.com/imroc/req"

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
func (user *userReq) login() (err error) {
	r := user.r

	sign := getSign(false, map[string]string{
		"deviceid":     user.devID,
		"imeiid":       user.imei,
		"phone":        user.phone,
		"pwd":          user.pwd,
		"account_type": accountType,
	})

	body := req.Param{
		"deviceid":     user.devID,
		"imeiid":       user.imei,
		"phone":        user.phone,
		"pwd":          user.pwd,
		"account_type": "4",
		"sign":         sign,
	}

	resp, err := r.Post(apiLoginURL, headers, body)
	if err != nil {
		return err
	}

	var v loginResp
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return errLogin
	}

	user.userInfo = &v.Data

	return
}

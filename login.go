package main

import (
	"fmt"
	"github.com/imroc/req"
)

type userInfo struct {
	Userid      string `json:"userid"`
	Phone       string `json:"phone"`
	AccountType string `json:"account_type"`
	BindPwd     int    `json:"bind_pwd"`
	NickName    string `json:"nickname"`
}

//map[iRet:0 sMsg:ok data:map[userid:<user id> phone:<phone> account_type:4 bind_pwd:1 nickname:<phone>]]
type logResp struct {
	Ret  int      `json:"iRet"`
	Mesg string   `json:"sMesg"`
	Data userInfo `json:"data"`
}

/*
login
*/
func login(r *req.Req, phone, pwd, dev, imei, sign string) (sessionid, userid string) {
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
		fmt.Println(err)
		return "", ""
	}

	for _, v := range resp.Response().Cookies() {
		switch v.Name {
		case "sessionid":
			sessionid = v.Value
		case "userid":
			userid = v.Value
		}
	}

	var v logResp
	if err := resp.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if v.Ret != 0 {
		fmt.Println("login fail")
		return
	}

	fmt.Println(v.Data, "\n")

	return
}

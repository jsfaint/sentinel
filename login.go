package main

import (
	"fmt"
	"github.com/imroc/req"
)

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

	var v map[string]interface{}
	if err := resp.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if !checkStatus(v) {
		fmt.Println("login fail")
		return
	}

	fmt.Println(v, "\n")

	return
}

package main

import (
	"fmt"
	"github.com/imroc/req"
	"math/big"
)

/*
login
*/
func login(phone, pwd, dev, imei, sign string) (sessionid, userid string) {
	//POST query parameter
	param := req.Param{
		"deviceid":     dev,
		"imeiid":       imei,
		"phone":        phone,
		"pwd":          pwd,
		"account_type": "4",
		"sign":         sign,
	}

	r, err := req.Post(apiLoginUrl, headers, param)
	if err != nil {
		fmt.Println(err)
		return "", ""
	}

	for _, v := range r.Response().Cookies() {
		switch v.Name {
		case "sessionid":
			sessionid = v.Value
		case "userid":
			userid = v.Value
		}
	}

	var v map[string]interface{}
	if err := r.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if checkStatus(v) {
		fmt.Println("login success")
	} else {
		fmt.Println("login fail")
	}

	fmt.Println(v, "\n")

	return
}

func checkStatus(v map[string]interface{}) bool {
	ret, ok := v["iRet"]
	if !ok {
		return false
	}

	num := big.NewFloat(ret.(float64))
	return (num.Cmp(big.NewFloat(0)) == 0)
}

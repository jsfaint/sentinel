package main

import (
	"fmt"
	"github.com/imroc/req"
	"math/big"
)

/*
checkStatus checks the iRet status
*/
func checkStatus(v map[string]interface{}) bool {
	ret, ok := v["iRet"]
	if !ok {
		return false
	}

	num := big.NewFloat(ret.(float64))
	return (num.Cmp(big.NewFloat(0)) == 0)
}

/*
checkRegister check if the phone number is already registered
*/
func checkRegister(r *req.Req, phone string) {
	fmt.Println("Phone:", phone)
	param := req.Param{
		"phone": phone,
	}

	resp, err := r.Post(apiCheckAccountURL, headers, param)
	if err != nil {
		fmt.Println(err)
		return
	}

	var v map[string]interface{}

	if err := resp.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if !checkStatus(v) {
		fmt.Println("checkRegister fail")
		return
	}

	fmt.Println(v, "\n")
}

func newReq() *req.Req {
	return req.New()
}

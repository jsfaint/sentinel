package main

import (
	"fmt"
	"github.com/imroc/req"
)

/*
{
  "iRet":0,
  "sMsg":"ok",
  "register":1
}
*/

type registResp struct {
	respHead
	Register int `json:"register"`
}

/*
checkRegister check if the phone number is already registered
*/
func checkRegister(r *req.Req, phone, sign string) (err error) {
	fmt.Println("Phone:", phone)
	param := req.Param{
		"phone": phone,
		"sign":  sign,
	}

	resp, err := r.Post(apiCheckAccountURL, headers, param)
	if err != nil {
		return err
	}

	var v registResp

	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return ERR_CHECK_REGISTERED
	}

	fmt.Println(v, "\n")

	return nil
}

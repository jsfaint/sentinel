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

type respRegister struct {
	respHead
	Register int `json:"register"`
}

/*
checkRegister check if the phone number is already registered
*/
func checkRegister(r *req.Req, phone, sign string) (err error) {
	fmt.Println("Phone:", phone)
	body := req.Param{
		"phone": phone,
		"sign":  sign,
	}

	resp, err := r.Post(apiCheckAccountURL, headers, body)
	if err != nil {
		return err
	}

	var v respRegister

	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return errCheckRegistered
	}

	fmt.Println(v, "\n")

	return nil
}

package main

import (
	"fmt"
	"github.com/imroc/req"
)

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

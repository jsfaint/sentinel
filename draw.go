package main

import (
	"fmt"

	"github.com/imroc/req"
)

type respDraw struct {
	respHead
}

func (user *userReq) withDraw() (err error) {
	r := user.r

	sign := getSign(false, map[string]string{
		"gasType":    "2",
		"drawWkb":    user.accountInfo.Balance,
		"appversion": appVersion,
	})

	body := req.Param{
		"gasType":    "2",
		"drawWkb":    user.accountInfo.Balance,
		"appversion": appVersion,
		"sign":       sign,
	}

	resp, err := r.Post(apiWKBDrawURL, body, headers)
	if err != nil {
		return err
	}

	fmt.Println(resp.Dump())

	return
}

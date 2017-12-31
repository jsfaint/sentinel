package main

import (
	"github.com/imroc/req"
)

var (
	headers = req.Header{
		"user-agent":    "Mozilla/5.0",
		"cache-control": "no-cache",
	}
)

type session struct {
	sessionid string
	userid    string
}

type userReq struct {
	phone string
	pwd   string
	devID string
	imei  string

	//Request
	r *req.Req

	//Data from response
	userInfo     *userInfo
	accountInfo  *accountInfo
	incomeInfo   *incomeInfo
	outcomeInfo  *outcomeInfo
	activateInfo *activateInfo
}

func newUser(phone, pass string) *userReq {
	return &userReq{
		phone: phone,
		pwd:   getPWD(pass),
		devID: getDevID(phone),
		imei:  getIMEI(phone),

		r: req.New(),
	}
}

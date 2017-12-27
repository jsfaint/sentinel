package main

import (
	"github.com/imroc/req"
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
	sign  string

	//Request
	r *req.Req
	session

	//Data from response
	userInfo     *userInfo
	accountInfo  *accountInfo
	incomeInfo   *incomeInfo
	outcomeInfo  *outcomeInfo
	activateInfo *activateInfo
}

func newUser(phone, pass string) *userReq {
	devid := getDevID(phone)
	pwd := getPWD(pass)
	imei := getIMEI(phone)
	sign := getSign(map[string]string{
		"deviceid":     devid,
		"imeiid":       imei,
		"phone":        phone,
		"pwd":          pwd,
		"account_type": accountType,
	})

	return &userReq{
		phone: phone,
		pwd:   getPWD(pass),
		devID: getDevID(phone),
		imei:  getIMEI(phone),
		sign:  sign,

		r: req.New(),
	}
}

package main

import (
	"fmt"
	"github.com/imroc/req"
)

var (
	headers = req.Header{
		"user-agent":    "Mozilla/5.0",
		"cache-control": "no-cache",
	}
)

type session struct {
	sessionID string
	userID    string
}

type userReq struct {
	phone string
	pwd   string
	devID string
	imei  string

	//Request
	r *req.Req
	session

	//Data from response
	userInfo     *userInfo
	accountInfo  *accountInfo
	incomeInfo   *incomeInfo
	outcomeInfo  *outcomeInfo
	activateInfo *activateInfo
	peers        *peerList
	partitions   *partitionList
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

func (user *userReq) refresh() (err error) {
	valid, err := user.validSession()
	if err != nil {
		return
	}

	if !valid {
		if err = user.login(); err != nil {
			return err
		}
	}

	if err = user.listPeerInfo(); err != nil {
		fmt.Println(err)
		return
	}

	if err = user.getUSBInfo(); err != nil {
		fmt.Println(err)
		return
	}

	if err = user.getActivate(); err != nil {
		fmt.Println(err)
		return
	}

	if err = user.getAccountInfo(); err != nil {
		fmt.Println(err)
		return
	}

	if err = user.getIncome(); err != nil {
		fmt.Println(err)
		return
	}

	if err = user.getOutcome(); err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (user *userReq) summary() (message string, err error) {

	return
}

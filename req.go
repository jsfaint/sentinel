package main

import (
	"bytes"
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

func (user *userReq) summary() string {
	var b bytes.Buffer

	//Short the variable
	account := user.accountInfo
	activate := user.activateInfo
	income := user.incomeInfo
	outcome := user.outcomeInfo
	peer := user.peers.Devices[0]
	partitions := user.partitions

	b.WriteString(fmt.Sprintf("电话: %s 钱包: %s\n", user.phone, account.Addr))
	b.WriteString(fmt.Sprintf("设备名: %s SN: %s\n", peer.DeviceName, peer.DeviceSn))
	b.WriteString(fmt.Sprintf("公网IP: %s 局域网IP: %s\n", peer.IP, peer.LanIP))
	b.WriteString(fmt.Sprintf("激活天数: %d 总收益: %.3f 已提币: %s\n", activate.ActivateDays, income.TotalIncome, outcome.TotalOutcome))
	b.WriteString(fmt.Sprintf("昨日收益: %.3f 可提余额: %s\n", activate.YesWKB, account.Balance))
	for _, p := range partitions.Partitions {
		b.WriteString(fmt.Sprintf("%s 容量: %sG/%sG", p.Label, p.Used, p.Capacity))
	}

	return b.String()
}

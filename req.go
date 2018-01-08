package main

import (
	"bytes"
	"fmt"
	"github.com/imroc/req"
	log "gopkg.in/clog.v1"
	"strconv"
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

type userData struct {
	//Data from response
	userInfo     *userInfo
	accountInfo  *accountInfo
	incomeInfo   *incomeInfo
	outcomeInfo  *outcomeInfo
	activateInfo *activateInfo
	peers        *peerList
	partitions   *partitionList
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
	userData
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

func (user *userReq) refresh(all bool) (err error) {
	valid, err := user.validSession()
	if err != nil {
		log.Info("%v", err)
	}

	if !valid {
		if err = user.login(); err != nil {
			return err
		}
	}

	if err = user.listPeerInfo(); err != nil {
		log.Error(0, "user.listPeerInfo() returns error %v", err)
	}

	if !all {
		return
	}

	phone := user.userInfo.Phone

	if user.peers.Devices[0].Status != "exception" {
		if err = user.getUSBInfo(); err != nil {
			log.Error(0, "user.getUSBInfo() returns error %v", err)
		}
	} else {
		log.Error(0, "%s disk status is exception %v", phone)
	}

	if err = user.getActivate(); err != nil {
		log.Error(0, "%s getActivate fail %v", phone, err)
	}

	if err = user.getAccountInfo(); err != nil {
		log.Error(0, "%s getAccountInfo fail %v", phone, err)
	}

	if err = user.getIncome(); err != nil {
		log.Error(0, "%s getIncome fail %v", phone, err)
	}

	if err = user.getOutcome(); err != nil {
		log.Error(0, "%s getOutcome fail %v", phone, err)
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

	b.WriteString(fmt.Sprintf("## %s\n", user.phone))
	b.WriteString(fmt.Sprintf("设备名: %s SN: %s  \n", peer.DeviceName, peer.DeviceSn))
	b.WriteString(fmt.Sprintf("公网IP: %s 局域网IP: %s  \n", peer.IP, peer.LanIP))
	b.WriteString(fmt.Sprintf("钱包地址: %s  \n", account.Addr))
	b.WriteString(fmt.Sprintf("激活天数: %d 总收益: %.3f 已提币: %s  \n", activate.ActivateDays, income.TotalIncome, outcome.TotalOutcome))
	b.WriteString(fmt.Sprintf("昨日收益: %.3f 可提余额: %s  \n", activate.YesWKB, account.Balance))
	for _, p := range partitions.Partitions {
		used, _ := strconv.ParseFloat(p.Used, 64)
		caps, _ := strconv.ParseFloat(p.Capacity, 64)

		used /= 1024 * 1024 * 1024
		caps /= 1024 * 1024 * 1024

		b.WriteString(fmt.Sprintf("%s 容量: %.2fG/%.2fG  \n", p.Label, used, caps))
	}

	return b.String()
}

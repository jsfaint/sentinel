package sentinel

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

//UserData defines all the user data which get from serve
type UserData struct {
	//Data from response
	UserInfo     *userInfo
	AccountInfo  *accountInfo
	IncomeInfo   *incomeInfo
	OutcomeInfo  *outcomeInfo
	ActivateInfo *activateInfo
	Peers        *peerList
	Partitions   *partitionList
}

//UserReq includes some info for request and response
type UserReq struct {
	phone string
	pwd   string
	devID string
	imei  string

	//Request
	r *req.Req
	session

	//Data from response
	UserData
}

//NewUser returns *UserReq
func NewUser(phone, pass string) *UserReq {
	return &UserReq{
		phone: phone,
		pwd:   getPWD(pass),
		devID: getDevID(phone),
		imei:  getIMEI(phone),

		r: req.New(),
	}
}

//Refresh will get all data from API
func (user *UserReq) Refresh(all bool) (err error) {
	valid, err := user.validSession()
	if err != nil {
		log.Info("%v", err)
	}

	if !valid {
		log.Info("%s Re-login", user.phone)
		user.r = req.New()
		if err = user.Login(); err != nil {
			return err
		}
	}

	if err = user.listPeerInfo(); err != nil {
		log.Error(0, "user.listPeerInfo() returns error %v", err)
		return err
	}

	if !all {
		return
	}

	phone := user.UserInfo.Phone

	if user.Peers.Devices[0].Status != "exception" {
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

//Summary will send information via servchan
func (user *UserReq) Summary() string {
	var b bytes.Buffer

	//Short the variable
	account := user.AccountInfo
	activate := user.ActivateInfo
	income := user.IncomeInfo
	outcome := user.OutcomeInfo
	peer := user.Peers.Devices[0]
	partitions := user.Partitions

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

//Phone returns user phone number
func (user *UserReq) Phone() string {
	return user.phone
}

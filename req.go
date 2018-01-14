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
	userInfo     *UserInfo
	accountInfo  *AccountInfo
	incomeInfo   *IncomeInfo
	outcomeInfo  *OutcomeInfo
	activateInfo *ActivateInfo
	peers        *Peers
	partitions   *Partitions
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

//Summary will send information via servchan
func (user *UserReq) Summary() string {
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

//Phone returns user phone number
func (user UserReq) Phone() string {
	return user.phone
}

//GetUserInfo provides API to get user information which returns struct UserInfo
func (user UserReq) GetUserInfo() UserInfo {
	if user.userInfo != nil {
		return *user.userInfo
	}
	return UserInfo{}
}

//GetAccountInfo provides API to get account information which returns struct AccountInfo
func (user UserReq) GetAccountInfo() AccountInfo {
	if user.accountInfo != nil {
		return *user.accountInfo
	}
	return AccountInfo{}
}

//GetIncome provides API to get incoming history which returns struct IncomeInfo
func (user UserReq) GetIncome() IncomeInfo {
	if user.incomeInfo != nil {
		return *user.incomeInfo
	}
	return IncomeInfo{}
}

//GetOutcome provides API to get outcoming history which returns struct OutcomeInfo
func (user UserReq) GetOutcome() OutcomeInfo {
	if user.outcomeInfo != nil {
		return *user.outcomeInfo
	}
	return OutcomeInfo{}
}

//GetPeers provides API to get peers information which returns struct Peers
func (user UserReq) GetPeers() Peers {
	if user.peers != nil {
		return *user.peers
	}
	return Peers{}
}

//GetActivateInfo provides API to get activation information which returns struct ActivateInfo
func (user UserReq) GetActivateInfo() ActivateInfo {
	if user.activateInfo != nil {
		return *user.activateInfo
	}
	return ActivateInfo{}
}

//GetPartitions provides API to parition information which returns struct Partitions
func (user UserReq) GetPartitions() Partitions {
	if user.partitions != nil {
		return *user.partitions
	}
	return Partitions{}
}

package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/imroc/req"
)

/*
{
    "rtn": 0,
    "result": [
        0,
        {
            "devices": [
                {
                    "bind_time": 1513126965,
                    "features": {
                        "onecloud_coin": 1513126984,
                        "miner": 0
                    },
                    "ip_info": {
                        "province": "",
                        "country": "",
                        "isp": "",
                        "city": ""
                    },
                    "ip": "",
                    "coturn_online": 1,
                    "paused": true,
                    "dcdn_upnp_message": "",
                    "device_sn": "",
                    "exception_message": "",
                    "account_name": "",
                    "dcdn_clients": [],
                    "hibernated": false,
                    "imported": "",
                    "exception_name": "",
                    "hardware_model": "",
                    "mac_address": "",
                    "status": "",
                    "lan_ip": "",
                    "account_type": "",
                    "account_id": "",
                    "upgradeable": false,
                    "dcdn_download_speed": 0,
                    "dcdn_upload_speed": 0,
                    "disk_quota": 10995116277760,
                    "peerid": "",
                    "licence": "",
                    "dcdn_id": "",
                    "system_version": "",
                    "device_id": "",
                    "system_name": "",
                    "dcdn_upnp_status": "",
                    "product_id": 412,
                    "device_name": "",
                    "schedule_hours": [
                        {
                            "to": 24,
                            "from": 0,
                            "params": {},
                            "type": "automatic"
                        }
                    ]
                }
            ]
        }
    ]
}
*/

type respListPeer struct {
	respHead2
	Result []interface{} `json:"result"`
}

type peerList struct {
	Devices []peerInfo `json:"devices"`
}

type peerInfo struct {
	BindTime int `json:"bind_time"`
	Features struct {
		OnecloudCoin int `json:"onecloud_coin"`
		Miner        int `json:"miner"`
	} `json:"features"`
	IPInfo struct {
		Province string `json:"province"`
		Country  string `json:"country"`
		Isp      string `json:"isp"`
		City     string `json:"city"`
	} `json:"ip_info"`
	IP                string        `json:"ip"`
	CoturnOnline      int           `json:"coturn_online"`
	Paused            bool          `json:"paused"`
	DcdnUpnpMessage   string        `json:"dcdn_upnp_message"`
	DeviceSn          string        `json:"device_sn"`
	ExceptionMessage  string        `json:"exception_message"`
	AccountName       string        `json:"account_name"`
	DcdnClients       []interface{} `json:"dcdn_clients"`
	Hibernated        bool          `json:"hibernated"`
	Imported          string        `json:"imported"`
	ExceptionName     string        `json:"exception_name"`
	HardwareModel     string        `json:"hardware_model"`
	MacAddress        string        `json:"mac_address"`
	Status            string        `json:"status"`
	LanIP             string        `json:"lan_ip"`
	AccountType       string        `json:"account_type"`
	AccountID         string        `json:"account_id"`
	Upgradeable       bool          `json:"upgradeable"`
	DcdnDownloadSpeed int           `json:"dcdn_download_speed"`
	DcdnUploadSpeed   int           `json:"dcdn_upload_speed"`
	DiskQuota         int64         `json:"disk_quota"`
	Peerid            string        `json:"peerid"`
	Licence           string        `json:"licence"`
	DcdnID            string        `json:"dcdn_id"`
	SystemVersion     string        `json:"system_version"`
	DeviceID          string        `json:"device_id"`
	SystemName        string        `json:"system_name"`
	DcdnUpnpStatus    string        `json:"dcdn_upnp_status"`
	ProductID         int           `json:"product_id"`
	DeviceName        string        `json:"device_name"`
	ScheduleHours     []struct {
		To     int `json:"to"`
		From   int `json:"from"`
		Params struct {
		} `json:"params"`
		Type string `json:"type"`
	} `json:"schedule_hours"`
}

/*
Get peer info
*/
func (user *userReq) listPeerInfo() (err error) {
	r := user.r

	sign := getSign(true, map[string]string{
		"appversion": wkAppVersion,
		"v":          "2",
		"ct":         "1",
	}, user.sessionID)

	query := req.QueryParam{
		"appversion": wkAppVersion,
		"v":          "2",
		"ct":         "1",
		"sign":       sign,
	}

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:  "origin",
			Value: "1",
		},
		&http.Cookie{
			Name:  "path",
			Value: "/",
		},
	}

	resp, err := r.Get(apiListPeerURL, query, cookies, headers)
	if err != nil {
		return
	}

	var v respListPeer

	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return v
	}

	if len(v.Result) != 2 {
		return errors.New("Invalid response result")
	}

	//Convert interface to json string and then convert to peerList
	var b []byte
	if b, err = json.Marshal(v.Result[1]); err != nil {
		return err
	}

	var list peerList
	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}

	user.peers = &list

	return
}

func (peer peerInfo) String(phone string) string {

	switch peer.Status {
	case "online":
		return phone + " " + "设备已离线"
	case "offline":
		return phone + " " + "设备已离线"
	case "exception":
		return phone + " " + "硬盘异常"

	}

	return ""
}

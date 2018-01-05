package main

import (
	log "gopkg.in/clog.v1"

	"errors"
	"github.com/imroc/req"
	"math/big"
)

/*
{
  "iRet": 0,
  "sMsg": "",
  "data": {
    "addr": "",
    "drawMsg": "",
    "drawRet": 0,
    "drawTime": "",
    "gasCost": "0.01",
    "gasType": 2,
    "originBalance": ""
  }
}
*/

type respDraw struct {
	respHead
	Data drawInfo `json:"data"`
}

type drawInfo struct {
	Addr          string `json:"addr"`
	DrawMsg       string `json:"draw_msg"`
	DrawRet       int    `json:"draw_ret"`
	DrawTime      string `json:"draw_time"`
	GasCost       string `json:"gas_cost"`
	GasType       int    `json:"gas_type"`
	OriginBalance string `json:"origin_balance"`
}

func balanceIsZero(s string) bool {
	balance, _, err := big.ParseFloat(s, 10, 64, big.ToZero)
	if err != nil {
		return true
	}

	if balance.Sign() == 0 {
		return true
	}

	return false
}

func (user *userReq) withDraw() (err error) {
	r := user.r

	if user.accountInfo == nil {
		return errors.New(user.phone + " account Info doesn't exist")
	}

	if balanceIsZero(user.accountInfo.Balance) {
		return errors.New("Drawed already")
	}

	sign := getSign(false, map[string]string{
		"gasType":    "2",
		"drawWkb":    user.accountInfo.Balance,
		"appversion": wkAppVersion,
	}, user.sessionID)

	body := req.Param{
		"gasType":    "2",
		"drawWkb":    user.accountInfo.Balance,
		"appversion": wkAppVersion,
		"sign":       sign,
	}

	resp, err := r.Post(apiWKBDrawURL, body, headers)
	if err != nil {
		return err
	}

	log.Info("%s", resp.Dump())

	var v respDraw
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	log.Info("%v", v)

	if !v.success() {
		return v
	}

	if !v.Data.success() {
		return v.Data
	}

	return
}

func (draw drawInfo) success() bool {
	if draw.DrawRet == 0 {
		return true
	}

	return false
}

func (draw drawInfo) Error() string {
	return draw.DrawMsg
}

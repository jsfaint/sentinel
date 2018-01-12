package sentinel

import (
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
	DrawMsg       string `json:"drawMsg"`
	DrawRet       int    `json:"drawRet"`
	DrawTime      string `json:"drawTime"`
	GasCost       string `json:"gasCost"`
	GasType       int    `json:"gasType"`
	OriginBalance string `json:"originBalance"`
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

//WithDraw request a coin drawing
func (user *UserReq) WithDraw() (err error) {
	r := user.r

	if user.AccountInfo == nil {
		return errors.New(user.phone + " account Info doesn't exist")
	}

	if balanceIsZero(user.AccountInfo.Balance) {
		return errors.New("Drawed already")
	}

	sign := getSign(false, map[string]string{
		"gasType":    "2",
		"drawWkb":    user.AccountInfo.Balance,
		"appversion": wkAppVersion,
	}, user.sessionID)

	body := req.Param{
		"gasType":    "2",
		"drawWkb":    user.AccountInfo.Balance,
		"appversion": wkAppVersion,
		"sign":       sign,
	}

	resp, err := r.Post(apiWKBDrawURL, body, headers)
	if err != nil {
		return err
	}

	var v respDraw
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

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

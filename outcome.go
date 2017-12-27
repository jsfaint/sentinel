package main

import "net/http"

/*
{
  "iRet": 0,
  "sMsg": "",
  "data": {
    "totalOutcome": "",
    "nextPage": -1,
    "outcomeArr": [
      {
        "date": "1970-01-01 00:00:00",
        "originBalance": 1.12345678,
        "gasCost": 0.01,
        "status": 2,
        "gasType": 0,
        "updateTime": "",
        "addr": ""
      },
      {
        "date": "1970-01-01 00:00:00",
        "originBalance": 1.12345678,
        "gasCost": 0.01,
        "status": 2,
        "gasType": 0,
        "updateTime": "",
        "addr": ""
      }
    ]
  }
}
*/

type outcomeResp struct {
	respHead
	Data outcomeInfo `json:"data"`
}

type outcomeInfo struct {
	TotalOutcome string         `json:"totalOutcome"`
	NextPage     int            `json:"nextPage"`
	OutcomeArr   []dailyOutcome `json:"outcomeArr"`
}

type dailyOutcome struct {
	Date          string  `json:"date"`
	OriginBalance float64 `json:"originBalance"`
	GasCost       float64 `json:"gasCost"`
	Status        int     `json:"status"`
	GasType       int     `json:"gasType"`
	UpdateTime    string  `json:"updateTime"`
	Addr          string  `json:"addr"`
}

/*
Get Outcome history
*/
func (user *userReq) getOutcome() (err error) {
	r := user.r

	cookies := []*http.Cookie{
		&http.Cookie{
			Name:  "sessionid",
			Value: user.sessionid,
		},
		&http.Cookie{
			Name:  "userid",
			Value: user.userid,
		},
		&http.Cookie{
			Name:  "origin",
			Value: "1",
		},
	}

	resp, err := r.Post(apiOutcomeURL, cookies, headers)
	if err != nil {
		return err
	}

	var v outcomeResp

	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return errOutcome
	}

	user.outcomeInfo = &v.Data

	return
}

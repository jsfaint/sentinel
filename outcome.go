package main

import (
	"fmt"
	"github.com/imroc/req"
	"net/http"
)

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
	Data outcome `json:"data"`
}

type outcome struct {
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

func getOutcome(r *req.Req, sessionid, userid string) (err error) {
	cookies := []*http.Cookie{
		&http.Cookie{
			Name:  "sessionid",
			Value: sessionid,
		},
		&http.Cookie{
			Name:  "userid",
			Value: userid,
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
		return ERR_OUTCOME
	}

	fmt.Println(v, "\n")

	return nil
}

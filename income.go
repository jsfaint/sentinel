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
    "totalIncome": 1.23456789,
    "nextPage": -1,
    "incomeArr": [
      { "date": "19700101", "num": "1.09098742" },
      { "date": "19700101", "num": "1.68708031" },
      { "date": "19700101", "num": "1.28912181" },
      { "date": "19700101", "num": "1.41720727" },
      { "date": "19700101", "num": "1.06814516" },
      { "date": "19700101", "num": "1.13231244" },
      { "date": "19700101", "num": "1.2342931" },
      { "date": "19700101", "num": "1.61914638" },
      { "date": "19700101", "num": "1.72968532" },
      { "date": "19700101", "num": "1.62592952" },
      { "date": "19700101", "num": "1.73835741" },
      { "date": "19700101", "num": "1.04103502" },
      { "date": "19700101", "num": "1.82745612" },
      { "date": "19700101", "num": "1.41449516" }
    ]
  }
}
*/

type incomeResp struct {
	respHead
	Data income `json:"data"`
}

type income struct {
	TotalIncome float64       `json:"totalIncome"`
	NextPage    int           `json:"nextPage"`
	IncomeArr   []dailyIncome `json:"incomeArr"`
}

type dailyIncome struct {
	Date string `json:"date"`
	Num  string `json:"num"`
}

func getIncome(r *req.Req, sessionid, userid, sign string) (err error) {
	param := req.Param{
		"page":       "0",
		"appversion": appVersion,
		"sign":       sign,
	}

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

	resp, err := r.Post(apiIncomeURL, param, cookies, headers)
	if err != nil {
		return
	}

	var v incomeResp

	if err := resp.ToJSON(&v); err != nil {
		fmt.Println(err)
		return err
	}

	if !v.success() {
		return fmt.Errorf("Get income fail")
	}

	fmt.Println(v, "\n")

	return nil
}

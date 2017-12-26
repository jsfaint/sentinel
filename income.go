package main

import (
	"fmt"
	"github.com/imroc/req"
	"net/http"
)

func income(sessionid, userid, sign string) {
	param := req.Param{
		"page":       "0",
		"appversion": "1.4.8",
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

	r, err := req.Post(apiIncomeURL, param, cookies, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	var v map[string]interface{}

	if err := r.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if checkStatus(v) {
		fmt.Println("Get income success")
	} else {
		fmt.Println("Get income fail")
	}

	fmt.Println(v, "\n")
}

package main

import (
	"fmt"
	"github.com/imroc/req"
	"net/http"
)

func outcome(sessionid, userid string) {
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

	r, err := req.Post(apiIncomeURL, cookies, headers)
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
		fmt.Println("Get outcome success")
	} else {
		fmt.Println("Get outcome fail")
	}

	fmt.Println(v, "\n")
}

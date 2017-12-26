package main

import (
	"fmt"
	"github.com/imroc/req"
	"net/http"
)

func outcome(r *req.Req, sessionid, userid string) {
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
		fmt.Println(err)
		return
	}

	var v map[string]interface{}

	if err := resp.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if !checkStatus(v) {
		fmt.Println("Get outcome fail")
		return
	}

	fmt.Println(v, "\n")
}

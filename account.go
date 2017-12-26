package main

import (
	"fmt"
	"github.com/imroc/req"
	"net/http"
)

/*
getAccountInfo
*/
func getAccountInfo(r *req.Req, sessionid, userid string) {
	//POST query parameter
	param := req.Param{
		"appversion": "1.4.8",
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

	resp, err := r.Post(apiAccountInfoURL, headers, param, cookies)
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
		fmt.Println("getAccountInfo fail")
		return
	}

	fmt.Println(v, "\n")
}

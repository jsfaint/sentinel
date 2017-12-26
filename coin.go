package main

import (
	"fmt"
	"github.com/imroc/req"
)

/*
getCoinInfo
*/
func getCoinInfo() {
	r, err := req.Post(apiCoinInfoURL, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	var v map[string]interface{}
	if err := r.ToJSON(&v); err != nil {
		fmt.Println(err)
		return
	}

	if !checkStatus(v) {
		fmt.Println("getCoinInfo fail")
	}

	fmt.Println(v, "\n")
}

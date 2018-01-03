package main

import (
	"github.com/imroc/req"
)

type respSession struct {
	respHead
}

func (user *userReq) validSession() (valid bool, err error) {
	r := user.r

	sign := getSign(true, map[string]string{}, user.sessionID)

	query := req.QueryParam{
		"appversion": wkAppVersion,
	}

	body := req.Param{
		"sign": sign,
	}

	resp, err := r.Post(apiSessionURL, query, body, headers)
	if err != nil {
		return false, err
	}

	var v respSession
	if err := resp.ToJSON(&v); err != nil {
		return false, err
	}

	if !v.success() {
		return false, v
	}

	return true, nil
}

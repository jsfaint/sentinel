package sentinel

import (
	"github.com/imroc/req"
)

const (
	servchanURL = "http://sc.ftqq.com/"
)

var servchanToken string

/*
{"errno":0,"errmsg":"success","dataset":"done"}
*/

type respSend struct {
	Errno   int    `json:"errno"`
	ErrMsg  string `json:"err_msg"`
	Dataset string `json:"dataset"`
}

//Send given info via serchan
func Send(title string, content string) (err error) {
	body := req.Param{
		"text": title,
		"desp": content,
	}

	url := servchanURL + servchanToken + ".send"

	resp, err := req.Post(url, body)
	if err != nil {
		return
	}

	var v respSend
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return v
	}

	return
}

func (s respSend) success() bool {
	return (s.Errno == 0)
}

func (s respSend) Error() string {
	return s.ErrMsg
}

//SetToken set servchan token
func SetToken(token string) {
	servchanToken = token
}

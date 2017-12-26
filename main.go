package main

import (
	"fmt"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Walk through the configs, support multiple account
	for _, u := range cfg.Accounts {
		phone := u.Phone
		pwd := getPWD(u.Pass)

		if phone == "" || pwd == "" {
			continue
		}

		dev := getDevID(phone)
		imei := getIMEI(phone)

		//Get sign
		sign := getSign(map[string]string{
			"deviceid":     dev,
			"imeiid":       imei,
			"phone":        phone,
			"pwd":          pwd,
			"account_type": "4",
		})

		println("登录")
		sessionid, userid := login(phone, pwd, dev, imei, sign)

		println("收益记录")
		income(sessionid, userid, sign)

		println("提币记录")
		outcome(sessionid, userid)
	}
}

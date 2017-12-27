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

		//skip if phone or pwd is null
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

		r := newReq()

		println("登录")
		sessionid, userid, err := login(r, phone, pwd, dev, imei, sign)
		if err != nil {
			fmt.Println(err)
			return
		}

		println("账号信息")
		if err = getAccountInfo(r, sessionid, userid); err != nil {
			fmt.Println(err)
			return
		}

		println("收益记录")
		if err = getIncome(r, sessionid, userid, sign); err != nil {
			fmt.Println(err)
			return
		}

		println("提币记录")
		outcome(r, sessionid, userid)
	}
}

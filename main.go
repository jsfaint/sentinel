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

	if c, err := getCoinInfo(); err != nil {
		fmt.Println(err)
	} else {
		_ = c //FIXME: make compiler happy
	}

	//Walk through the configs, support multiple account
	for _, u := range cfg.Accounts {
		//skip if phone or pwd is null
		if u.Phone == "" || u.Pass == "" {
			continue
		}

		user := newUser(u.Phone, u.Pass)

		println("登录")
		if err = user.login(); err != nil {
			fmt.Println(err)
			return
		}

		//if err = user.listPeerInfo(); err != nil {
		//fmt.Println(err)
		//return
		//}

		println("账号信息")
		if err = user.getAccountInfo(); err != nil {
			fmt.Println(err)
			return
		}

		println("收益记录")
		if err = user.getIncome(); err != nil {
			fmt.Println(err)
			return
		}

		println("提币记录")
		if err = user.getOutcome(); err != nil {
			fmt.Println(err)
			return
		}

		user.withDraw()
	}
}

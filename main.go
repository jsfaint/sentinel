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

	//Set servchan token
	setServtoken(cfg.Token)

	//Walk through the configs, support multiple account
	var users []*userReq
	for _, u := range cfg.Accounts {
		//skip if phone or pwd is null
		if u.Phone == "" || u.Pass == "" {
			continue
		}

		user := newUser(u.Phone, u.Pass)

		if user != nil {
			users = append(users, user)
		}

		println("登陆")
		if err = user.login(); err != nil {
			fmt.Println(err)
			continue
		}
	}

	println("获取玩客币信息")
	if c, err := getCoinInfo(); err != nil {
		fmt.Println(err)
	} else {
		_ = c //FIXME: make compiler happy
	}

	for _, user := range users {
		if user.userInfo == nil {
			fmt.Println(user.phone, "log in failed, please check the username or password")
			continue
		}

		println("验证session有效性")
		if _, err := user.validSession(); err != nil {
			fmt.Println(err)
			return
		}

		println("获取节点列表")
		if err := user.listPeerInfo(); err != nil {
			fmt.Println(err)
			return
		}

		println("获取磁盘信息")
		if err := user.getUSBInfo(); err != nil {
			fmt.Println(err)
			return
		}

		println("获取激活信息")
		if err := user.getActivate(); err != nil {
			fmt.Println(err)
			return
		}

		println("账号信息")
		if err := user.getAccountInfo(); err != nil {
			fmt.Println(err)
			return
		}

		println("收益记录")
		if err := user.getIncome(); err != nil {
			fmt.Println(err)
			return
		}

		println("提币记录")
		if err := user.getOutcome(); err != nil {
			fmt.Println(err)
			return
		}

		println("提币")
		if err := user.withDraw(); err != nil {
			fmt.Println(err)
		}
	}
}

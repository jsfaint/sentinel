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
	}

	println("获取玩客币信息")
	go func() {
		c, err := getCoinInfo()
		if err != nil {
			fmt.Println(err)
		}

		c.dump()
	}()

	login(users)

	refresh(users)

	println("summary")
	summary(users)

	println("checkStatus")
	checkStatus(users)
}

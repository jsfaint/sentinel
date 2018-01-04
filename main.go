package main

import (
	"flag"
	"fmt"
)

const (
	name        = "sentinel"
	description = "A sentinel for WanKeYun"
	version     = "v0.1"
	mail        = "jsfaint@gmail.com"
)

type options struct {
	summary *bool
	check   *bool
	draw    *bool
}

var opt options

func init() {
	fmt.Println(name, version, mail)
	fmt.Println(description)

	opt.summary = flag.Bool("summary", false, "Show summary")
	opt.check = flag.Bool("check", false, "Check device status")
	opt.draw = flag.Bool("draw", false, "Draw the incoming coins")

	flag.Parse()
}

func main() {
	//Walk through the configs, support multiple account
	var users []*userReq
	for _, u := range cfg.Accounts {
		//skip if phone or pwd is null
		if u.Phone == "" || u.Pass == "" {
			continue
		}

		user := newUser(u.Phone, u.Pass)

		if err := user.login(); err != nil {
			fmt.Println(err)
			continue
		}

		users = append(users, user)
	}

	//If some parameter were given, run the specfied task else run cron job
	if *opt.summary || *opt.check || *opt.draw {
		//Refresh all data
		refresh(users)

		if *opt.summary {
			summary(users)
		}

		if *opt.check {
			checkStatus(users)
		}

		if *opt.draw {
			withDraw(users)
		}

		return
	}

	if err := startCrontable(users); err != nil {
		return
	}
	defer stopCrontable()

	select {}
}

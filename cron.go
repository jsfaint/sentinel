package main

import (
	"fmt"
	"github.com/robfig/cron"
	log "gopkg.in/clog.v1"
	"time"
)

/*
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
*/

var c *cron.Cron

func startCrontable(users []*userReq) (err error) {
	c = cron.New()

	var tz int
	if time.Now().Location().String() == "UTC" {
		tz = 8
	} else {
		tz = 0
	}

	summaryTime := fmt.Sprintf("0 0 %d * * *", 9-tz)
	drawTime := fmt.Sprintf("0 10 %d * * %d", 9-tz, cfg.DrawDay)

	//Refresh data and check status every 30s
	if err = c.AddFunc("@every 30s", func() {
		checkStatus(users)
	}); err != nil {
		log.Error(0, "Add checkStatus() fail %v", err)
	}

	//Show summary every morning 9:00
	if err = c.AddFunc(summaryTime, func() {
		refresh(users)
		summary(users)
	}); err != nil {
		log.Error(0, "Add summary() fail%v", err)
	}

	//Draw at 9:10 of the setting day
	if err = c.AddFunc(drawTime, func() {
		refresh(users)
		withDraw(users)
	}); err != nil {
		log.Error(0, "Add withDraw() fail %v", err)
	}

	c.Start()

	return
}

func stopCrontable() {
	c.Stop()
}

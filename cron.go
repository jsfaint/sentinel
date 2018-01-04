package main

import (
	"fmt"
	"github.com/robfig/cron"
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

	//Refresh data and check status every 30s
	if err = c.AddFunc("30 * * * * *", func() {
		refresh(users)
		checkStatus(users)
	}); err != nil {
		return err
	}

	//Show summary every morning 9:00
	if err = c.AddFunc("0 0 9 * * *", func() {
		summary(users)
	}); err != nil {
		return err
	}

	//Draw at 9:10 of the setting day
	if err = c.AddFunc(fmt.Sprintf("0 10 9 * * %d", cfg.DrawDay), func() {
		withDraw(users)
	}); err != nil {
		return err
	}

	c.Start()

	return
}

func stopCrontable() {
	c.Stop()
}

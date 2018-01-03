package main

import (
	"fmt"
	"github.com/robfig/cron"
)

var c *cron.Cron

func startCrontable(users []*userReq) (err error) {
	c = cron.New()

	//Refresh data and check status every 30s
	if err = c.AddFunc("@every 30s", func() {
		refresh(users)
		checkStatus(users)
	}); err != nil {
		return err
	}

	//Show summary every morning 9:00
	if err = c.AddFunc("00 09 * * *", func() {
		summary(users)
	}); err != nil {
		return err
	}

	//Draw at 9:10 of the setting day
	if err = c.AddFunc(fmt.Sprintf("10 09 * * %d", cfg.DrawDay), func() {
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

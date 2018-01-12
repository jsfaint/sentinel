package main

import (
	"bytes"
	"fmt"
	log "gopkg.in/clog.v1"
	"sentinel"
	"sync"
)

//Login all account in goroutine
func login(users []*sentinel.UserReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *sentinel.UserReq) {
			defer done.Done()

			if err := u.Login(); err != nil {
				log.Error(0, "u.login() return errors %v", err)
			}

		}(u)
	}

	done.Wait()
}

//Refresh the data to avoid out-of-date
func refresh(users []*sentinel.UserReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *sentinel.UserReq) {
			defer done.Done()

			phone := u.UserData.UserInfo.Phone

			if err := u.Refresh(true); err != nil {
				log.Error(0, "%s refresh() returns error %v", phone, err)
			}

			log.Trace("%s refreshed", phone)
		}(u)
	}

	done.Wait()
}

//With draw the coin
func withDraw(users []*sentinel.UserReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *sentinel.UserReq) {
			defer done.Done()

			if err := u.WithDraw(); err != nil {
				log.Error(0, "u.withDraw() returns error %v", err)
			}
		}(u)
	}

	done.Wait()
}

//Show summary of yesterday
func summary(users []*sentinel.UserReq) {
	var b bytes.Buffer

	b.WriteString(incomeAverage(users))
	b.WriteString("\n")

	for _, u := range users {
		b.WriteString(u.Summary())
		b.WriteString("\n")
	}

	log.Info("%s", b.String())

	if err := sentinel.Send("玩客哨兵每日播报", b.String()); err != nil {
		log.Error(0, "Send summary info to servchan fail %v", err)
	}
}

//Check online status and send alarming
func checkStatus(users []*sentinel.UserReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *sentinel.UserReq) {
			defer done.Done()

			if err := u.Refresh(false); err != nil {
				log.Error(0, "u.refresh() returns error %v", err)
				return
			}

			for _, v := range u.Peers.Devices {
				if v.Status == "online" {
					continue
				}

				phone := u.Phone()

				t, c := v.Message(phone)

				if err := sentinel.Send(t, c); err != nil {
					log.Error(0, "send() returns error %s %v", phone, err)
				}
			}
		}(u)
	}

	done.Wait()
}

//Get Average income of yesterday
func incomeAverage(users []*sentinel.UserReq) string {
	var total float64
	var b bytes.Buffer

	for _, u := range users {
		total += u.ActivateInfo.YesWKB
	}

	b.WriteString(fmt.Sprintf("共%d台机器 总收益 %.3f链克  \n", len(users), total))
	b.WriteString(fmt.Sprintf("平均%.3f 链克/台  \n", total/float64(len(users))))

	return b.String()
}

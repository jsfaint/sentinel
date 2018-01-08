package main

import (
	"bytes"
	"fmt"
	log "gopkg.in/clog.v1"
	"sync"
)

//Login all account in goroutine
func login(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *userReq) {
			defer done.Done()

			if err := u.login(); err != nil {
				log.Error(0, "u.login() return errors %v", err)
			}

		}(u)
	}

	done.Wait()
}

//Refresh the data to avoid out-of-date
func refresh(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *userReq) {
			defer done.Done()

			if err := u.refresh(true); err != nil {
				log.Error(0, "%s refresh() returns error %v", u.userData.userInfo.Phone, err)
			}

			log.Trace("%s refreshed", u.userData.userInfo.Phone)
		}(u)
	}

	done.Wait()
}

//With draw the coin
func withDraw(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *userReq) {
			defer done.Done()

			if err := u.withDraw(); err != nil {
				log.Error(0, "u.withDraw() returns error %v", err)
			}
		}(u)
	}

	done.Wait()
}

//Show summary of yesterday
func summary(users []*userReq) {
	var b bytes.Buffer

	b.WriteString(incomeAverage(users))
	b.WriteString("\n")

	for _, u := range users {
		b.WriteString(u.summary())
		b.WriteString("\n")
	}

	log.Info("%s", b.String())

	if err := send("玩客哨兵每日播报", b.String()); err != nil {
		log.Error(0, "Send summary info to servchan fail %v", err)
	}
}

//Check online status and send alarming
func checkStatus(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *userReq) {
			defer done.Done()

			if err := u.refresh(false); err != nil {
				log.Error(0, "u.refresh() returns error %v", err)
				return
			}

			for _, v := range u.peers.Devices {
				if v.Status == "online" {
					continue
				}

				t, c := v.Message(u.phone)

				if err := send(t, c); err != nil {
					log.Error(0, "send() returns error %s %v", u.phone, err)
				}
			}
		}(u)
	}

	done.Wait()
}

//Get Average income of yesterday
func incomeAverage(users []*userReq) string {
	var total float64
	var b bytes.Buffer

	for _, u := range users {
		total += u.activateInfo.YesWKB
	}

	b.WriteString(fmt.Sprintf("共%d台机器 总收益 %.3f链克  \n", len(users), total))
	b.WriteString(fmt.Sprintf("平均%.3f 链克/台  \n", total/float64(len(users))))

	return b.String()
}

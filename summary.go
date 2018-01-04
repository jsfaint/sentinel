package main

import (
	"bytes"
	"fmt"
	"sync"
)

var oldUsers []*userReq

//Login all account in goroutine
func login(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *userReq) {
			if err := u.login(); err != nil {
				fmt.Println(err)
			}

			done.Done()
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
			if err := u.refresh(); err != nil {
				fmt.Println(err)
			}

			done.Done()
		}(u)
	}

	done.Wait()

	oldUsers = users
}

//With draw the coin
func withDraw(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func(u *userReq) {
			if err := u.withDraw(); err != nil {
				fmt.Println(err)
			}

			done.Done()
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

	if err := send("玩客哨兵每日播报", b.String()); err != nil {
		fmt.Println(err)
	}
}

//Check online status and send alarming
func checkStatus(users []*userReq) {
	if len(oldUsers) == 0 {
		return
	}

	var done sync.WaitGroup

	for i, u := range users {
		done.Add(1)
		go func(u *userReq, old *userReq) {
			for i, v := range u.peers.Devices {
				status := old.peers.Devices[i].Status
				if v.Status == status {
					continue
				}

				t, c := v.Message(u.phone)

				if err := send(t, c); err != nil {
					fmt.Println(u.phone, err)
				}
			}

			done.Done()
		}(u, oldUsers[i])
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

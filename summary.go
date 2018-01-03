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
		go func() {
			if err := u.login(); err != nil {
				fmt.Println(err)
			}

			done.Done()
		}()
	}

	done.Wait()
}

//Refresh the data to avoid out-of-date
func refresh(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func() {
			if err := u.refresh(); err != nil {
				fmt.Println(err)
			}

			done.Done()
		}()
	}

	done.Wait()

	oldUsers = users
}

//With draw the coin
func withDraw(users []*userReq) {
	var done sync.WaitGroup

	for _, u := range users {
		done.Add(1)
		go func() {
			if err := u.withDraw(); err != nil {
				fmt.Println(err)
			}

			done.Done()
		}()
	}

	done.Wait()
}

//Show summary of yesterday
func summary(users []*userReq) {
	var b bytes.Buffer

	b.WriteString(incomeAverage(users))

	for _, u := range users {
		b.WriteString(u.summary())
	}

	if err := send("summary", b.String()); err != nil {
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

				s := v.String(u.phone)

				if err := send(s, s); err != nil {
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

	b.WriteString(fmt.Sprintf("%.3f台机器\n", len(users)))
	b.WriteString(fmt.Sprintf("%.3f 币/台", total/float64(len(users))))

	return b.String()
}

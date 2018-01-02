package main

import (
	"fmt"
	"strings"
	"sync"
)

var oldUsers []*userReq

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

func summary(users []*userReq) {
	var sum []string

	for _, u := range users {
		message, err := u.summary()
		if err != nil {
			fmt.Println(err)
		}

		sum = append(sum, message)
	}

	for _, v := range sum {
		fmt.Println(v)
	}

	if err := send("summary", strings.Join(sum, "\n")); err != nil {
		fmt.Println(err)
	}
}

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

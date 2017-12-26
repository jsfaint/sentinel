package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type user struct {
	Phone string `json:"phone"`
	Pass  string `json:"pass"`
}

type config struct {
	Token string `json:"token"`
	Users []user `json:"users"`
}

func Config() (cfg config, err error) {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
		return config{}, err
	}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		fmt.Println(err.Error())
		return config{}, err
	}

	fmt.Println(cfg)

	return
}

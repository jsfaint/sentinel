package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var cfg config

type account struct {
	Phone string `json:"phone"`
	Pass  string `json:"pass"`
}

type config struct {
	Token    string       `json:"token"`
	DrawDay  time.Weekday `json:"draw_day"`
	Accounts []account    `json:"accounts"`
}

func getConfig() (cfg config, err error) {
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

	return
}

func init() {
	var err error
	cfg, err = getConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if len(cfg.Accounts) == 0 {
		log.Fatalln("Please add at least one account in config.json")
	}

	if cfg.Token == "" {
		log.Fatalln("Empty servchan token, please set token in the config.json")
	}

	//Set servchan token
	setServtoken(cfg.Token)
}

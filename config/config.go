package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

//Config defines servchan token, draw day and account info
type Config struct {
	Token    string       `json:"token"`
	DrawDay  time.Weekday `json:"draw_day"`
	Accounts []struct {
		Phone string `json:"phone"`
		Pass  string `json:"pass"`
	} `json:"accounts"`
}

//NewConfig get configuration from config.json
func NewConfig() (cfg *Config, err error) {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return
}

//Check config data
func (cfg *Config) Check() error {
	if len(cfg.Accounts) == 0 {
		return errors.New("Please add at least one account in config.json")
	}

	if cfg.DrawDay < 1 || cfg.DrawDay > 5 {
		return errors.New("Invalid draw day, must be in the range of 1~5")
	}

	if cfg.Token == "" {
		return errors.New("Empty servchan token, please set token in the config.json")
	}

	return nil
}

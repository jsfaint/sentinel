package main

import (
	"errors"
)

var (
	errAccountInfo     = errors.New("get AccountInfo fail")
	errActivateInfo    = errors.New("Get activateInfo fail")
	errCheckRegistered = errors.New("check register status fail")
	errCoinInfo        = errors.New("Get coin info fail")
	errIncomeInfo      = errors.New("Get income fail")
	errLogin           = errors.New("Login Fail")
	errOutcome         = errors.New("Get outcome fail")
)

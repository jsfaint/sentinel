package main

import (
	"errors"
)

var (
	ERR_ACCOUNT_INFO     = errors.New("getAccountInfo fail")
	ERR_CHECK_REGISTERED = errors.New("checkRegister fail")
	ERR_OUTCOME          = errors.New("Get outcome fail")
	ERR_LOGIN            = errors.New("Login Fail")
	ERR_INCOME           = errors.New("Get income fail")
)

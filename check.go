package main

import (
	"math/big"
)

type respHead struct {
	Ret int    `json:"iRet"`
	Msg string `json:"sMesg"`
}

func (h *respHead) success() bool {
	if h.Ret != 0 {
		return false
	} else {
		return true
	}
}

/*
checkStatus checks the iRet status
*/
func checkStatus(v map[string]interface{}) bool {
	ret, ok := v["iRet"]
	if !ok {
		return false
	}

	num := big.NewFloat(ret.(float64))
	return (num.Cmp(big.NewFloat(0)) == 0)
}

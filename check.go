package main

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

package main

type respHead struct {
	Ret int    `json:"iRet"`
	Msg string `json:"sMesg"`
}

func (h *respHead) success() bool {
	return (h.Ret == 0)
}

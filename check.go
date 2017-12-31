package main

type respHead struct {
	Ret int    `json:"iRet"`
	Msg string `json:"sMesg"`
}

type respHead2 struct {
	Rtn int `json:"rtn"`
}

func (h respHead) success() bool {
	return (h.Ret == 0)
}

func (h respHead2) success() bool {
	return (h.Rtn == 0)
}

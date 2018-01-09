package main

type respHead struct {
	Ret int    `json:"iRet"`
	Msg string `json:"sMsg"`
}

type respHead2 struct {
	Rtn int    `json:"rtn"`
	Msg string `json:"msg"`
}

func (h respHead) success() bool {
	return (h.Ret == 0)
}

func (h respHead2) success() bool {
	return (h.Rtn == 0)
}

func (h respHead) Error() string {
	return h.Msg
}

func (h respHead2) Error() string {
	return h.Msg
}

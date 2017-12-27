package main

import (
	"fmt"
	"github.com/imroc/req"
)

/*
{
  "iRet": 0,
  "sMsg": "ok",
  "data": {
    "date": "2017\u5e7412\u670826\u65e5",
    "block_num": 434540,
    "wkb_num": 123843900,
    "average_onlinetime": 22,
    "average_bandwidth": 15,
    "average_disk": 1202,
    "topN_wkb": [
      { "user_id": 2039570, "wkb": 7, "phone": "139****8408" },
      { "user_id": 1781662, "wkb": 7, "phone": "181****2501" },
      { "user_id": 2907769, "wkb": 7, "phone": "171****5752" },
      { "user_id": 1013183, "wkb": 7, "phone": "133****8233" },
      { "user_id": 1008767, "wkb": 7, "phone": "181****2980" }
    ],
    "topN_bandwidth": [
      {
        "bandwidth": 668,
        "user_id": 2907769,
        "wkb": 7.89127854,
        "phone": "171****5752"
      },
      {
        "bandwidth": 607,
        "user_id": 2576434,
        "wkb": 7.42185316,
        "phone": "178****3189"
      },
      {
        "bandwidth": 568,
        "user_id": 1410485,
        "wkb": 7.24551283,
        "phone": "188****8204"
      },
      {
        "bandwidth": 568,
        "user_id": 2039570,
        "wkb": 7.89127854,
        "phone": "139****8408"
      },
      {
        "bandwidth": 547,
        "user_id": 1781662,
        "wkb": 7.89127854,
        "phone": "181****2501"
      }
    ],
    "topN_disk": [
      {
        "disk": 70355,
        "user_id": 1035124,
        "wkb": 5.66532284,
        "phone": "159****7269"
      },
      {
        "disk": 27912,
        "user_id": 1406474,
        "wkb": 5.41684095,
        "phone": "186****0822"
      },
      {
        "disk": 21034,
        "user_id": 1935411,
        "wkb": 2.30804269,
        "phone": "138****0007"
      },
      {
        "disk": 20931,
        "user_id": 2154685,
        "wkb": 2.33294204,
        "phone": "138****8559"
      },
      {
        "disk": 20789,
        "user_id": 1238763,
        "wkb": 5.7304069,
        "phone": "135****8006"
      }
    ]
  }
}
*/

type coinResp struct {
	respHead
	Data coin `json:"data"`
}

type coin struct {
	Date              string          `json:"date"`
	BlockNum          uint64          `json:"block_num"`
	WKBNum            uint64          `json:"wkb_num"`
	AverageOnlinetime uint8           `json:"average_onlinetime"`
	AverageBandwidth  uint32          `json:"average_bandwidth"`
	AverageDisk       uint64          `json:"average_disk"`
	TopNWKB           []topNWKB       `json:"topN_wkb"`
	TopNBandwidth     []topNBandwidth `json:"topN_bandwidth"`
}

type topInfo struct {
	UserID uint64  `json:"user_id"`
	WKB    float64 `json:"wkb"`
	Phone  string  `json:"phone"`
}

type topNWKB struct {
	topInfo
}

type topNBandwidth struct {
	topInfo
	Bandwidth uint32 `json:"bandwidth"`
}

type topNDisk struct {
	topInfo
	Disk uint64 `json:"disk"`
}

/*
getCoinInfo
*/
func getCoinInfo() (c *coin, err error) {
	r, err := req.Post(apiCoinInfoURL, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	var v coinResp
	if err := r.ToJSON(&v); err != nil {
		return nil, err
	}

	if !v.success() {
		return nil, errCoinInfo
	}

	c = &v.Data

	return
}

func (c *coin) dump() {
	fmt.Println(c)
}

package sentinel

import (
	"encoding/json"
	"errors"

	"github.com/imroc/req"
)

/*
{
  "msg":"success",
  "rtn":0,
  "result":[
    0,
    {
      "partitions":[
        {
          "used":"1479147520",
          "capacity":"999659470848",
          "disk_id":0,
          "label":"/dev/sda1",
          "path":"/media/sda1",
          "unique":23,
          "id":1
        }
      ]
    }
  ]
}
*/

type respUSBInfo struct {
	respHead2
	Result []interface{} `json:"result"`
}

type partitionList struct {
	Partitions []partitionInfo `json:"partitions"`
}

type partitionInfo struct {
	Used     string `json:"used"`
	Capacity string `json:"capacity"`
	DiskID   int    `json:"disk_id"`
	Label    string `json:"label"`
	Path     string `json:"path"`
	Unique   int    `json:"unique"`
	ID       int    `json:"id"`
}

func (user *UserReq) getUSBInfo() (err error) {
	r := user.r

	devID := user.Peers.Devices[0].DeviceID

	sign := getSign(true, map[string]string{
		"appversion": wkAppVersion,
		"v":          "1",
		"ct":         "1",
		"deviceid":   devID,
	}, user.sessionID)

	query := req.QueryParam{
		"appversion": wkAppVersion,
		"v":          "1",
		"ct":         "1",
		"deviceid":   devID,
		"sign":       sign,
	}

	resp, err := r.Get(apiUSBInfoURL, query)
	if err != nil {
		return err
	}

	var v respUSBInfo
	if err := resp.ToJSON(&v); err != nil {
		return err
	}

	if !v.success() {
		return v
	}

	if len(v.Result) != 2 {
		return errors.New("Invalid response result")
	}

	//Convert interface to json string and then convert to partitionList
	var b []byte
	if b, err = json.Marshal(v.Result[1]); err != nil {
		return err
	}

	var list partitionList
	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}

	user.Partitions = &list

	return err
}

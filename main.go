package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	//"github.com/imroc/req"
	"math"
	"strconv"
	"strings"
)

const (
	appVersion     = "1.4.5"
	apiAccountUrl  = "https://account.onethingpcs.com"
	apiControlUrl  = "https://control.onethingpcs.com"
	apiRemoteDlUrl = "https://control.remotedl.onethingpcs.com"
)

var (
	headers = map[string]string{"user-agent": "Mozilla/5.0"}
)

func md5LowerString(s string) string {
	m := md5.New()
	b := m.Sum([]byte(s))
	str := hex.EncodeToString(b)

	return strings.ToLower(str)
}

func deviceID(phone string) string {
	s := md5LowerString(phone)

	//Convert to upper case
	return strings.ToUpper(s)
}

//Get password
func password(text string) string {
	s := md5LowerString(text)
	return md5LowerString(s)
}

func imei(phone string) string {
	num, err := strconv.ParseFloat(phone, 10)
	if err != nil {
		fmt.Println("Convert string to int fail")
		return ""
	}

	s := fmt.Sprint("%f", math.Pow(num, 2))

	return s[0:14]
}

func main() {
	fmt.Println("vim-go")
}

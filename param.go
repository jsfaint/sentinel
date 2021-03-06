package sentinel

import (
	log "gopkg.in/clog.v1"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	wkAppVersion = "1.4.13"
	accountType  = "4"
)

/*
getDevID generate device id
*/
func getDevID(phone string) string {
	s := md5LowerString(phone)

	//Convert to upper case
	return strings.ToUpper(s[0:16])
}

/*
getPWD Get getPWD string
*/
func getPWD(text string) string {
	s := md5LowerString(text)

	str := s[0:2] + string(s[8]) + s[3:8] + string(s[2]) + s[9:17] +
		string(s[27]) + s[18:27] + string(s[17]) + s[28:]

	return md5LowerString(str)
}

/*
getSign calculate the sign via some config
*/
func getSign(isGet bool, body map[string]string, session string) string {
	var list []string

	//Generate list
	for k, v := range body {
		list = append(list, k+"="+v)
	}

	//Sort
	sort.Strings(list)

	//Join
	s := strings.Join(list, "&") + "&key=" + session

	return md5LowerString(s)
}

/*
getIMEI generate IMEI via phone number, it's not a real imem number
*/
func getIMEI(phone string) string {
	num, err := strconv.ParseFloat(phone, 64)
	if err != nil {
		log.Error(0, "Convert string to int fail")
		return ""
	}

	s := strconv.FormatFloat(math.Pow(num, 2), 'f', 6, 64)

	return s[0:14]
}

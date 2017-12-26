package main

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

/*
md5LowerString sum md5 hash code with lower case
*/
func md5LowerString(s string) string {
	b := md5.Sum([]byte(s))
	str := hex.EncodeToString(b[:])

	return strings.ToLower(str)
}

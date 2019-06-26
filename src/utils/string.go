package utils

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// MakeNString makes n-length string given unit-string
func MakeNString(n int, s string) string {
	resString := ""
	for i := 0; i < n; i++ {
		resString += s
	}
	return resString
}

// Escape escape cetrain string
func Escape(target string, query string) string {
	return strings.Replace(target, query, fmt.Sprintf("\\%s", query), -1)
}

// EncodeTag encode with base64 and then replace `/` and `+`
// into `a_a` and `b_b` respectively.
func EncodeTag(tag string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(tag))
	sEnc = strings.Replace(sEnc, "/", "a_a", -1)
	sEnc = strings.Replace(sEnc, "+", "b_b", -1)
	sEnc = strings.Replace(sEnc, "=", "c_c", -1)
	return sEnc
}

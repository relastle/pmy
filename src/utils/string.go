package utils

import (
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

package utils

// MakeNString makes n-length string given unit-string
func MakeNString(n int, s string) string {
	resString := ""
	for i := 0; i < n; i++ {
		resString += s
	}
	return resString
}

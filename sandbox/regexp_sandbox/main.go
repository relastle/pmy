package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`^cd +(?P<path>.+)$`)
	body := "cd q"
	fmt.Println(re.MatchString(body))
	fmt.Printf("%q\n", re.SubexpNames())
	fmt.Println(re.FindStringSubmatch(body))
}

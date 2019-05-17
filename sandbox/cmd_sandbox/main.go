package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func toLines(bs []byte) []string {
	strList := strings.Split(string(bs), "\n")
	return strList
}

func main() {
	name := "ls"
	args := []string{"-la"}
	out, err := exec.Command(name, args...).Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := toLines(out)
	for _, line := range lines {
		fmt.Println(line)
	}
}

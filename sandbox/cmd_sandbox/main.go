package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	pipeline "github.com/mattn/go-pipeline"
	shellwords "github.com/mattn/go-shellwords"
)

func toLines(bs []byte) []string {
	strList := strings.Split(string(bs), "\n")
	return strList
}

// pipeTest:
func pipeTest() {
	out, err := pipeline.Output(
		[]string{"ls", "/Users/hkonishi", "-alh"},
		[]string{"wc", "-l"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	// Output:
	// 1
}

// subprocessTest:
func subprocessTest() {
	shellwords.ParseBacktick = true
	p := shellwords.NewParser()
	p.ParseEnv = true
	args, err := p.Parse("${HOME} -alh | fzf")
	fmt.Println(args)
	out, err := exec.Command("ls", args...).Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := toLines(out)
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	pipeTest()
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/fatih/color"
)

type pmyTestCase struct {
	Lbuffer  string `json:"lbuffer"`
	Rbuffer  string `json:"rbuffer"`
	Expected string `json:"expected"`
}

func (c *pmyTestCase) testSelf(t *testing.T) (bool, error) {
	start := time.Now()
	out, err := exec.Command(
		"../shell/pmy_wrapper.zsh",
		c.Lbuffer,
		c.Rbuffer,
		"test",
	).Output()
	if err != nil {
		return false, err
	}

	res := strings.Replace(string(out), "\n", "", -1)
	if res == c.Expected {
		elapsed := time.Since(start)
		fmt.Printf(
			"[%v] pass; {res: %v, elapsed: %v}\n",
			color.GreenString("●"),
			res,
			elapsed,
		)
		return true, nil
	}
	fmt.Printf(
		"[%v] fail\nexpected: %v\nactual  : %v\n",
		color.RedString("✘"),
		c.Expected,
		res,
	)
	return false, nil
}

type pmyTestCases []*pmyTestCase

// TestIntegration conduct integration test
// - pmy core runner (Go part)
// - zsh-go interaction
func TestIntegration(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	os.Setenv(
		"PMY_RULE_PATH",
		fmt.Sprintf("%v/../test/rules", cwd),
	)
	os.Setenv(
		"PMY_SNIPPET_PATH",
		fmt.Sprintf("%v/../test/snippets", cwd),
	)
	jsonFile, err := os.Open(fmt.Sprintf("%v/../test/pmy_testcases.json", cwd))
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var cases pmyTestCases
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &cases)
	if err != nil {
		t.Error(err)
	}
	for _, c := range cases {
		if ok, err := c.testSelf(t); ok {
			continue
		} else if err == nil {
			t.Fail()
		} else {
			log.Fatal(err)
		}
	}
}

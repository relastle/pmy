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

	"github.com/fatih/color"
)

type pmyTestCase struct {
	Lbuffer  string `json:"lbuffer"`
	Rbuffer  string `json:"rbuffer"`
	Expected string `json:"expected"`
}

func (c *pmyTestCase) testSelf() (bool, error) {
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
		fmt.Printf(
			"[%v] pass; res: %v\n",
			color.GreenString("●"),
			res,
		)
		return true, nil
	}
	fmt.Printf(
		"[%v] fail\nexpectd: %v\nactual: %v\n",
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
	gopath := os.Getenv("GOPATH")
	os.Setenv(
		"PMY_RULE_PATH",
		fmt.Sprintf("%v/src/github.com/relastle/pmy/resources/test/pmy_rules_test.json", gopath),
	)
	jsonFile, err := os.Open("../resources/test/pmy_testcases.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var cases pmyTestCases
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &cases)
	for _, c := range cases {
		if ok, err := c.testSelf(); ok {
			continue
		} else if err == nil {
			t.Fail()
		} else {
			log.Fatal(err)
		}
	}
}

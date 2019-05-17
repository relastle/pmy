package pmy

import (
	"fmt"
	"log"
	"os/exec"
)

// ccheckJq checks whether `jq` command is available
func checkJq() bool {
	_, err := exec.LookPath("jq")
	if err != nil {
		return false
	}
	return true
}

// Run runs the main process of pmy
func Run(cfgPath string, bufferLeft string, bufferRight string) {

	// Load rules from config file
	rules, err := loadAllRules(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch rule using LBUFFER and RBUFFER
	var fetcher ruleFetcher
	fetcher = &ruleFetcherImpl{}
	resRules, err := fetcher.fetch(rules, bufferLeft, bufferRight)
	if err != nil {
		log.Fatal(err)
	}
	if len(resRules) == 0 {
		fmt.Print("")
		return
	}
	rule := resRules[0]
	out := newPmyOutFromRule(&rule)

	// Ouput result
	fmt.Println(out.serialize())
	return
}

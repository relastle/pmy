package pmy

import (
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

// Input represents input from zsh
type Input struct {
	BufferLeft  string
	BufferRight string
}

// Run runs the main process of pmy.
// It returns zsh statement, where resulting values will
// be passed into zsh variables.
func Run(cfgPath string, in Input) string {

	// Load rules from config file
	rules, err := loadAllRules(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch rule using LBUFFER and RBUFFER
	var fetcher ruleFetcher
	fetcher = &ruleFetcherImpl{}
	resRules, err := fetcher.fetch(
		rules,
		in.BufferLeft,
		in.BufferRight,
	)
	if err != nil {
		log.Fatal(err)
	}
	if len(resRules) == 0 {
		return ""
	}
	rule := resRules[0]
	pmyOut := newPmyOutFromRule(rule)

	// Ouput result
	// log.Print(out)
	return pmyOut.toShellVariables()
}

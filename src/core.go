package pmy

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

// checkJq checks whether `jq` command is available
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

func (i *Input) getCmdName() string {
	elms := strings.Split(i.BufferLeft, " ")
	if len(elms) == 0 {
		return ""
	}
	return elms[0]
}

// Run runs the main process of pmy.
// It returns zsh statement, where resulting values will
// be passed into zsh variables.
func Run(cfgPath string, in Input) string {
	// Load global rules from config file
	rules, err := loadAllRules(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Load command specific rules from config file
	if cmdName := in.getCmdName(); cmdName != "" {
		cmdCfgPath := filepath.Join(
			filepath.Dir(cfgPath),
			fmt.Sprintf("%v_%v", cmdName, filepath.Base(cfgPath)),
		)
		cmdRules, err := loadAllRules(cmdCfgPath)
		if err == nil {
			rules = append(cmdRules, rules...)
		}
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

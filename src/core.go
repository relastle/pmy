package pmy

import (
	"log"
	"sort"
	"strings"
	"time"
)

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
func Run(in Input) string {
	defer MeasureElapsedTime(time.Now(), "core.Run")

	// Load all rule files
	ruleFiles := GetAllRuleFiles()

	// Extract only applicable rule files
	ruleFilesToApply := []*RuleFile{}
	for _, ruleFile := range ruleFiles {
		if ruleFile.isApplicable(in.getCmdName()) {
			ruleFilesToApply = append(ruleFilesToApply, ruleFile)
		}
	}
	// Sort rule file by priority
	sort.SliceStable(
		ruleFilesToApply,
		func(i, j int) bool { return ruleFilesToApply[i].priority > ruleFilesToApply[j].priority },
	)
	rules := Rules{}
	for _, ruleFile := range ruleFilesToApply {
		_rules, err := ruleFile.loadRules()
		if err == nil {
			rules = append(rules, _rules...)
		} else {
			log.Fatal(err)
		}
	}

	// Fetch rule using LBUFFER and RBUFFER
	fetcher := &ruleFetcherImpl{}
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
	Out := newOutFromRule(rule)

	// Ouput result
	// log.Print(out)
	return Out.toShellVariables()
}

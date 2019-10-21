package pmy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/mattn/go-zglob"
	"gopkg.in/yaml.v2"
)

const (
	pmyRuleSuffixCommon = "pmy_rules.*"
	pmyRuleSuffixJSON   = "pmy_rules.json"
	pmyRuleSuffixYML    = "pmy_rules.yml"
	priorityGlobal      = 1
	priorityCmdSpecific = 2
)

// RuleFile represents one Rule Json file
// information
type RuleFile struct {
	Path     string
	Basename string
	priority int
}

func (rf RuleFile) loadRules() (Rules, error) {
	f, err := os.Open(rf.Path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var rules Rules
	byteValue, _ := ioutil.ReadAll(f)

	// Unmarshal using json or yml encoder
	if strings.HasSuffix(rf.Basename, "json") {
		err = json.Unmarshal(byteValue, &rules)
	} else if strings.HasSuffix(rf.Basename, "yml") {
		err = yaml.Unmarshal(byteValue, &rules)
	}
	if err != nil {
		return nil, err
	}
	return rules, nil
}

func (rf *RuleFile) isApplicable(cmd string) bool {
	if rf.Basename == pmyRuleSuffixJSON || rf.Basename == pmyRuleSuffixYML {
		rf.priority = priorityGlobal
		return true
	}
	if rf.Basename == fmt.Sprintf(
		"%s_%s",
		cmd,
		pmyRuleSuffixJSON,
	) || rf.Basename == fmt.Sprintf(
		"%s_%s",
		cmd,
		pmyRuleSuffixYML,
	) {
		rf.priority = priorityCmdSpecific
		return true
	}
	return false
}

// GetAllRuleFiles get all pmy rules json paths
// configured by environment variable
func GetAllRuleFiles() []*RuleFile {
	ruleRoots := strings.Split(RulePath, ":")
	ruleRoots = append(ruleRoots, defaultRulePath)

	res := []*RuleFile{}
	for _, ruleRoot := range ruleRoots {
		if ruleRoot == "" {
			continue
		}
		globPattern := fmt.Sprintf(
			`%v/**/*%v`,
			os.ExpandEnv(ruleRoot),
			pmyRuleSuffixCommon,
		)
		matches, err := zglob.Glob(globPattern)
		if err != nil {
			panic(err)
		}
		for _, rulePath := range matches {
			res = append(
				res,
				&RuleFile{
					Path:     rulePath,
					Basename: path.Base(rulePath),
				},
			)
		}

	}
	return res
}

// Rule is a struct representing one rule
type Rule struct {
	Name           string    `json:"name" yaml:"name"`
	Matcher        string    `json:"matcher" yaml:"matcher"`
	Description    string    `json:"description" yaml:"description"`
	RegexpLeft     string    `json:"regexpLeft" yaml:"regexp-left"`
	RegexpRight    string    `json:"regexpRight" yaml:"regexp-right"`
	CmdGroups      CmdGroups `json:"cmdGroups" yaml:"cmd-groups"`
	FuzzyFinderCmd string    `json:"fuzzyFinderCmd" yaml:"fuzzy-finder-cmd"`
	BufferLeft     string    `json:"bufferLeft" yaml:"buffer-left"`
	BufferRight    string    `json:"bufferRight" yaml:"buffer-right"`
	paramMap       map[string]string
}

// Rules represents slice of `Rule` struct
type Rules []*Rule

// match check if the query buffers(left and right) satisfies the concerned
// rule. if the rule regexp contains parametrized subgroups, this function expand
// the `Command` to `CommandExpanded`, where all parametrized variables were expanded.
func (rule *Rule) match(bufferLeft string, bufferRight string) (bool, error) {
	re, err := regexp.Compile(rule.RegexpLeft)
	if err != nil {
		return false, err
	}
	matches := re.FindStringSubmatch(bufferLeft)
	names := re.SubexpNames()
	if len(matches) != len(names) {
		return false, nil
	}
	paramMap := map[string]string{}
	for i, name := range names {
		if i != 0 && name != "" {
			paramMap[name] = matches[i]
		}
	}
	rule.BufferLeft = strings.Replace(
		rule.BufferLeft,
		"[]",
		bufferLeft,
		-1,
	)
	rule.BufferRight = strings.Replace(
		rule.BufferRight,
		"[]",
		bufferRight,
		-1,
	)
	rule.paramMap = paramMap
	return true, nil
}

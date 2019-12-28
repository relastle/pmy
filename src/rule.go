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
	// !!! will be deprecated !!!
	_pmyRuleSuffixJSON = "pmy_rules.json"
	_pmyRuleSuffixYML  = "pmy_rules.yml"
	_pmyRuleSuffixYAML = "pmy_rules.yaml"

	pmyRuleSuffixJSON = ".json"
	pmyRuleSuffixYML  = ".yml"
	pmyRuleSuffixYAML = ".yaml"

	priorityGlobal      = 1
	priorityCmdSpecific = 2
)

var (
	pmyRuleGlobalNames = []string{
		"_.json",
		"_.yml",
		"_.yaml",
	}
)

// RuleFile represents one Rule Json file
// information
type RuleFile struct {
	Path     string
	Basename string
	priority int
}

func (rf RuleFile) isGlobal() bool {
	for _, globalName := range pmyRuleGlobalNames {
		if rf.Basename == globalName {
			return true
		}
	}
	// TODO: deprecated global rules
	return (rf.Basename == _pmyRuleSuffixJSON ||
		rf.Basename == _pmyRuleSuffixYML ||
		rf.Basename == _pmyRuleSuffixYAML)
}

func (rf RuleFile) isCommandSpecific(cmd string) bool {
	if rf.Basename == fmt.Sprintf(
		"%s%s",
		cmd,
		pmyRuleSuffixJSON,
	) || rf.Basename == fmt.Sprintf(
		"%s%s",
		cmd,
		pmyRuleSuffixYML,
	) || rf.Basename == fmt.Sprintf(
		"%s%s",
		cmd,
		pmyRuleSuffixYAML,
	) {
		return true
	}

	// TODO: deprecated command specific rules
	if rf.Basename == fmt.Sprintf(
		"%s_%s",
		cmd,
		_pmyRuleSuffixJSON,
	) || rf.Basename == fmt.Sprintf(
		"%s_%s",
		cmd,
		_pmyRuleSuffixYML,
	) || rf.Basename == fmt.Sprintf(
		"%s_%s",
		cmd,
		_pmyRuleSuffixYAML,
	) {
		return true
	}
	return false
}

func (rf RuleFile) isJSON() bool {
	return strings.HasSuffix(rf.Basename, pmyRuleSuffixJSON)
}

func (rf RuleFile) isYAML() bool {
	return (strings.HasSuffix(rf.Basename, pmyRuleSuffixYML) ||
		strings.HasSuffix(rf.Basename, pmyRuleSuffixYAML))
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
	if rf.isJSON() {
		err = json.Unmarshal(byteValue, &rules)
	} else if rf.isYAML() {
		err = yaml.Unmarshal(byteValue, &rules)
	}
	if err != nil {
		return nil, err
	}
	return rules, nil
}

// setPriority set priority for the rule given a command
func (rf *RuleFile) setPriority(cmd string) bool {
	if rf.isGlobal() {
		rf.priority = priorityGlobal
		return true
	}

	if rf.isCommandSpecific(cmd) {
		rf.priority = priorityCmdSpecific
		return true

	}
	return false
}

// GetAllRuleFiles get all files under $PMY_RULE_PATH
// configured by environment variable
func GetAllRuleFiles() []*RuleFile {
	ruleRoots := strings.Split(RulePath, ":")
	ruleRoots = append(ruleRoots, defaultRulePath)

	res := []*RuleFile{}
	for _, ruleRoot := range ruleRoots {
		// expand environment variable
		ruleRoot = os.ExpandEnv(ruleRoot)
		if ruleRoot == "" {
			continue
		}
		globPattern := fmt.Sprintf(
			`%v/**/*`,
			ruleRoot,
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

package pmy

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type replaceMap map[string]string

// Rule is a struct representing one rule
type pmyRule struct {
	Name           string    `json:"name"`
	Matcher        string    `json:"matcher"`
	Description    string    `json:"description"`
	RegexpLeft     string    `json:"regexpLeft"`
	RegexpRight    string    `json:"regexpRight"`
	CmdGroups      CmdGroups `json:"cmdGroups"`
	FuzzyFinderCmd string    `json:"fuzzyFinderCmd"`
	BufferLeft     string    `json:"bufferLeft"`
	BufferRight    string    `json:"bufferRight"`
	paramMap       map[string]string
}

type pmyRules []*pmyRule

func loadAllRules(cfgPath string) (pmyRules, error) {
	jsonFile, err := os.Open(cfgPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	var rules pmyRules
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &rules)
	return rules, nil
}

// match check if the query buffers(left and right) satisfies the concerned
// rule. if the rule regexp contains parametrized subgroups, this function expand
// the `Command` to `CommandExpanded`, where all parametrized variables were expanded.
func (rule *pmyRule) match(bufferLeft string, bufferRight string) (bool, error) {
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

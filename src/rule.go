package pmy

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Rule is a struct representing one rule
type pmyRule struct {
	RegexpLeft  string `json:"regexpLeft"`
	RegexpRight string `json:"regexpRight"`
	Command     string `json:"command"`
}

type pmyRules []pmyRule

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

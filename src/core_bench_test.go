package pmy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// DumpDummyRulesJSON dumps arbitrary number of rules into given file path
func DumpDummyRulesJSON(resultPath string, ruleNum int, cmdGroupNum int) error {
	Rules := Rules{}
	for i := 0; i < ruleNum; i++ {
		cgs := CmdGroups{}
		for j := 0; j < cmdGroupNum; j++ {
			cg := &CmdGroup{
				Tag:   fmt.Sprintf("test%v", ruleNum),
				Stmt:  "find /Users/hkonishi/ -maxdepth 2",
				After: "awk '{print $1}'",
			}
			cgs = append(cgs, cg)
		}
		rule := &Rule{
			Name:        fmt.Sprintf("test%v", ruleNum),
			RegexpLeft:  ".*test.*",
			RegexpRight: "",
			CmdGroups:   cgs,
			BufferLeft:  "[]",
			BufferRight: "[]",
		}
		Rules = append(Rules, rule)
	}

	cgsJSON, _ := json.Marshal(Rules)
	err := ioutil.WriteFile(resultPath, cgsJSON, 0644)
	if err != nil {
		return err
	}
	return nil
}

// BenchmarkLoadLargeRules2 calculates time elapsed to load large rule file
// and fetch them all
func BenchmarkLoadLargeRules2(b *testing.B) {
	const ruleNum int = 10000
	const cmdGroupNum int = 1
	filePath := "./test_pmy_rules_large.json"
	DumpDummyRulesJSON(filePath, ruleNum, cmdGroupNum)
	b.ResetTimer()
	out := Run(
		filePath,
		Input{
			BufferLeft:  "git abcdef tes abcdef  ",
			BufferRight: "",
		},
	)
	b.StopTimer()
	os.Remove(filePath)
	if out != "" {
		log.Fatal("output is not empty")
	}
	return
}

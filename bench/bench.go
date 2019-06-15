package main

import (
	"fmt"
	"os"

	pmy "github.com/relastle/pmy/src"
)

func main() {
	const ruleNum int = 1
	const cmdGroupNum int = 1
	filePath := "./test_pmy_rules_large.json"
	pmy.DumpDummyRulesJSON(filePath, ruleNum, cmdGroupNum)
	outString := pmy.Run(
		filePath,
		pmy.Input{
			BufferLeft:  "atesta",
			BufferRight: "",
		},
	)
	fmt.Println(outString)
	os.Remove(filePath)
	return
}

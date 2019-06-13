package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	pmy "github.com/relastle/pmy/src"
)

// BenchmarkLoadandFetchAll calculates time elapsed to load test input rules
// and to fetch inputs to all of them
func BenchmarkLoadandFetchAll(b *testing.B) {
	inputJSONFile, err := os.Open("./resources/test/test_pmy_input_unmatch.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	defer inputJSONFile.Close()
	var inputs []pmy.Input
	b.ResetTimer()
	byteValue, _ := ioutil.ReadAll(inputJSONFile)
	json.Unmarshal(byteValue, &inputs)
	for _, input := range inputs {
		pmy.Run(
			"./resources/test/test_pmy_rules_large.json",
			input,
		)
	}
	return
}

// BenchmarkLoadLargeRules calculates time elapsed to load large rule file
func BenchmarkLoadLargeRules(b *testing.B) {
	pmy.Run(
		"./resources/test/test_pmy_rules_large.json",
		pmy.Input{
			BufferLeft:  "git abcdef add abcdef  ",
			BufferRight: "t aa e aa s aa t",
		},
	)
	return
}

package pmy

import (
	"fmt"
	"log"
)

// Run runs the main process of pmy
func Run(cfgPath string, bufferLeft string, bufferRight string) {
	rules, err := loadAllRules(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	var fetcher ruleFetcher
	fetcher = &ruleFetcherImpl{}
	resRules, err := fetcher.fetch(rules, bufferLeft, bufferRight)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resRules)
}

package pmy

import (
	"fmt"
	"log"
)

// Run runs the main process of pmy
func Run(cfgPath string) {
	rules, err := loadAllRules(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rules)
}

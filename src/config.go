package pmy

import (
	"log"
	"os"
)

const (
	pmyRulesPathVarName string = "PMY_RULE_PATH"
)

var (
	// PmyRulePath is a json path contining rules
	PmyRulePath string
)

func setConfig(
	target *string,
	varName string,
) {
	envVar, ok := os.LookupEnv(varName)
	if !ok {
		log.Fatalf("env var %v is not set", varName)
	}
	*target = envVar
}

// SetConfigs set all Pmy config variable from shell's
// environment variables.
func SetConfigs() {
	setConfig(&PmyRulePath, pmyRulesPathVarName)
}

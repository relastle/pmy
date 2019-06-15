package pmy

import (
	"log"
	"os"
)

const (
	pmyRulesPathVarName    string = "PMY_RULE_PATH"
	pmyTagDelimiterVarName string = "PMY_TAG_DELIMITER"
)

var (
	// PmyRulePath is a json path contining rules
	PmyRulePath string
	// PmyDelimiter defines delimiter string
	// that divide `tag` and one line of source
	PmyDelimiter string
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
	setConfig(&PmyDelimiter, pmyTagDelimiterVarName)
}

package pmy

import (
	"os"
)

const (
	defaultRulePath     string = "${HOME}/.pmy/rules"
	defaultSnippetPath  string = "${HOME}/.pmy/snippets"
	rulesPathVarName    string = "PMY_RULE_PATH"
	snippetPathVarName  string = "PMY_SNIPPET_PATH"
	tagDelimiterVarName string = "PMY_TAG_DELIMITER"
)

var (
	// RulePath is a json path contining rules
	RulePath string = ""
	// SnippetPath defines snippet root directry path
	SnippetPath string = ""
	// TagDelimiter defines delimiter string
	// that divide `tag` and one line of source
	TagDelimiter string = "\\t"
)

// setConfig set go varible from the given environment variable
func setConfig(
	target *string,
	varName string,
) {
	envVar, ok := os.LookupEnv(varName)
	if !ok {
		return
	}
	*target = envVar
}

// SetConfigs set all Pmy config variable from shell's
// environment variables.
func SetConfigs() {
	setConfig(&RulePath, rulesPathVarName)
	setConfig(&SnippetPath, snippetPathVarName)
	setConfig(&TagDelimiter, tagDelimiterVarName)
}

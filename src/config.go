package pmy

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-zglob"
)

const (
	defaultRulePath     string = "${HOME}/.pmy/rules"
	defaultSnippetPath  string = "${HOME}/.pmy/snippets"
	RulesPathVarName    string = "PMY_RULE_PATH"
	SnippetPathVarName  string = "PMY_SNIPPET_PATH"
	TagDelimiterVarName string = "PMY_TAG_DELIMITER"
)

var (
	// RulePath is a json path contining rules
	RulePath string
	// Delimiter defines delimiter string
	// that divide `tag` and one line of source
	Delimiter string
	// SnippetPath defines snippet root directry path
	SnippetPath string
)

// setConfig set go varible from the given environment variable
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
	setConfig(&RulePath, RulesPathVarName)
	setConfig(&Delimiter, TagDelimiterVarName)
	setConfig(&SnippetPath, SnippetPathVarName)
}

// GetAllRuleJSONPaths get all pmy rules json paths
// configured by environment variable
func GetAllRuleJSONPaths() []string {
	ruleRoots := []string{defaultRulePath}
	ruleRoots = append(ruleRoots, strings.Split(RulePath, ":")...)
	res := []string{}
	for _, ruleRoot := range ruleRoots {
		globPattern := fmt.Sprintf(`%v/**/*pmy_rules.json`, os.ExpandEnv(ruleRoot))
		matches, err := zglob.Glob(globPattern)
		if err != nil {
			panic(err)
		}
		res = append(res, matches...)
	}
	return res
}

// GetAllSnippetJSONPaths get all pmy rules json paths
// configured by environment variable
func GetAllSnippetJSONPaths() []string {
	snippetsRoots := []string{defaultSnippetPath}
	snippetsRoots = append(snippetsRoots, strings.Split(SnippetPath, ":")...)
	res := []string{}
	for _, snippetsRoot := range snippetsRoots {
		globPattern := fmt.Sprintf(`%v/**/*.txt`, os.ExpandEnv(snippetsRoot))
		matches, err := zglob.Glob(globPattern)
		if err != nil {
			panic(err)
		}
		res = append(res, matches...)
	}
	return res
}

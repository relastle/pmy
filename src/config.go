package pmy

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-zglob"
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
	// Delimiter defines delimiter string
	// that divide `tag` and one line of source
	Delimiter string = "\t"
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
	setConfig(&Delimiter, tagDelimiterVarName)
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

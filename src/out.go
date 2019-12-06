package pmy

import (
	"fmt"
	"strings"

	utils "github.com/relastle/pmy/src/utils"
)

const (
	shellBufferLeftVariableName     = "__pmy_out_buffer_left"
	shellBufferRightVariableName    = "__pmy_out_buffer_right"
	shellCommandVariableName        = "__pmy_out_command"
	shellAfterVariableName          = "__pmy_out_%s_after"
	shellFuzzyFinderCmdVariableName = "__pmy_out_fuzzy_finder_cmd"
	shellTagAllEmptyVariableName    = "__pmy_out_tag_all_empty"
	shellTagDelimiterVariableName   = "__pmy_out_tag_delimiter"
	shellErrorMessageVariableName   = "__pmy_out_error_message"
)

// Out represents Output of pmy against zsh routine.
// This struct has strings exported to shell, whose embedded
// variables are all expanded.
type Out struct {
	bufferLeft     string
	bufferRight    string
	cmdGroups      CmdGroups
	fuzzyFinderCmd string
	allEmptyTag    bool
	errorMessage   string
}

// newOutFromRule create new Out from rule
// which matches query and already has paramMap
func newOutFromRule(rule *Rule) Out {
	out := Out{}
	// pass resulting buffer informaiton
	out.bufferLeft = rule.BufferLeft
	out.bufferRight = rule.BufferRight
	// pass cmdGroups
	out.cmdGroups = rule.CmdGroups
	out.cmdGroups.alignTag()
	out.fuzzyFinderCmd = rule.FuzzyFinderCmd
	// expand all params
	out.expandAllParams(rule.paramMap)
	// expand magic command
	err := out.expandAllMagics()
	if err != nil {
		out.errorMessage += err.Error() + "\n"
	}
	// check if all tag is empty string
	out.allEmptyTag = out.cmdGroups.allEmpty()
	return out
}

// buildMainCommand builds main command that concatenate
// all results of given command groups
func (out *Out) buildMainCommand() string {
	res := ""
	for _, cg := range out.cmdGroups {
		// if there is no tag, simply execute command
		if out.allEmptyTag {
			res += fmt.Sprintf(
				"%v ;",
				cg.Stmt,
			)
		} else {
			res += fmt.Sprintf(
				"%v | awk '{printf \"%%s%v%%s\\\\n\", \"%v\", $0}' ; ",
				cg.Stmt,
				utils.EscapeBackslash(TagDelimiter),
				cg.tagAligned,
			)
		}
	}
	return res
}

// toShellVariables create zsh statement where Out's attributes are
// passed into shell variables
func (out *Out) toShellVariables() string {
	res := ""
	res += fmt.Sprintf("local %v=$'%v';", shellCommandVariableName, utils.Escape(out.buildMainCommand(), "'"))
	res += fmt.Sprintf("local %v=$'%v';", shellBufferLeftVariableName, utils.Escape(out.bufferLeft, "'"))
	res += fmt.Sprintf("local %v=$'%v';", shellBufferRightVariableName, utils.Escape(out.bufferRight, "'"))
	res += fmt.Sprintf("local %v=$'%v';", shellFuzzyFinderCmdVariableName, utils.Escape(out.fuzzyFinderCmd, "'"))
	res += fmt.Sprintf("local %v='%v';", shellTagDelimiterVariableName, TagDelimiter)
	res += fmt.Sprintf("local %v=$'%v';", shellErrorMessageVariableName, utils.Escape(out.errorMessage, "'"))
	if out.allEmptyTag {
		res += fmt.Sprintf("local %v=$'%v';", shellTagAllEmptyVariableName, "empty")
	}

	for _, cg := range out.cmdGroups {
		res += fmt.Sprintf(
			"local %v=$'%v';",
			fmt.Sprintf(shellAfterVariableName, utils.EncodeTag(cg.tagAligned)),
			utils.Escape(cg.After, "'"),
		)
	}
	return res
}

func expand(org string, paramMap map[string]string) string {
	res := org
	for name, value := range paramMap {
		res = strings.Replace(
			res,
			fmt.Sprintf("<%v>", name),
			value,
			-1,
		)
	}
	return res
}

func getToApply(snippetFiles []*SnippetFile, relpath string) (*SnippetFile, error) {
	for _, snippetFile := range snippetFiles {
		if snippetFile.isApplicable(relpath) {
			return snippetFile, nil
		}

	}
	return nil, fmt.Errorf("snippet file %v not found", relpath)
}

// expandAllMagics expands all magic commnad in Stmt
// written in `%hoge` format
func (out *Out) expandAllMagics() error {
	snippetFiles := GetAllSnippetFiles()
	for _, cg := range out.cmdGroups {
		if !strings.HasPrefix(cg.Stmt, "%") {
			continue
		}
		startIndex := strings.Index(cg.Stmt, "%") + 1      // inclusive
		length := strings.Index(cg.Stmt[startIndex:], "%") // exclusive
		snippetRelPath := cg.Stmt[startIndex : startIndex+length]
		snippetFileToApply, err := getToApply(snippetFiles, snippetRelPath)
		if err != nil {
			return err
		}

		cg.Stmt = strings.Replace(
			cg.Stmt,
			"%"+snippetRelPath+"%",
			fmt.Sprintf("cat %s", snippetFileToApply.Path),
			-1,
		)
	}
	return nil
}

// expandAllParams expands all params that refer to regexp parameters
func (out *Out) expandAllParams(paramMap map[string]string) {
	out.bufferLeft = expand(out.bufferLeft, paramMap)
	out.bufferRight = expand(out.bufferRight, paramMap)
	out.fuzzyFinderCmd = expand(out.fuzzyFinderCmd, paramMap)
	for _, cg := range out.cmdGroups {
		cg.Stmt = expand(cg.Stmt, paramMap)
	}
}

package pmy

import (
	"encoding/json"
	"fmt"
	"os"
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
)

type afterCmd struct {
	tag   string
	after string
}

// pmyOut represents Output of pmy against zsh routine.
// This struct has strings exported to shell, whose embedded
// variables are all expanded.
type pmyOut struct {
	bufferLeft     string
	bufferRight    string
	cmdGroups      CmdGroups
	sources        string
	fuzzyFinderCmd string
	allEmptyTag    bool
}

// newPmyOutFromRule create new pmyOut from rule
// which matches query and already has paramMap
func newPmyOutFromRule(rule *pmyRule) pmyOut {
	out := pmyOut{}
	// pass resulting buffer informaiton
	out.bufferLeft = rule.BufferLeft
	out.bufferRight = rule.BufferRight
	// pass cmdGroups
	out.cmdGroups = rule.CmdGroups
	out.cmdGroups.alignTag()
	out.fuzzyFinderCmd = rule.FuzzyFinderCmd
	// expand magic command
	out.expandAllMagics()
	// expand magic command
	out.expandAllParams(rule.paramMap)
	// check if all tag is empty string
	out.allEmptyTag = out.cmdGroups.allEmpty()
	return out
}

// buildMainCommand builds main command that concatenate
// all results of given command groups
func (out *pmyOut) buildMainCommand() string {
	res := ""
	pmyDelimiter := os.Getenv("PMY_TAG_DELIMITER")
	for _, cg := range out.cmdGroups {
		// if there is no tag, no need to use taggo
		if out.allEmptyTag {
			res += fmt.Sprintf(
				"%v ;",
				cg.Stmt,
			)
		} else {
			res += fmt.Sprintf(
				"%v | taggo -t '%v' -c '%v' --tag-delimiter '%v' ;",
				cg.Stmt,
				cg.tagAligned,
				cg.TagColor,
				pmyDelimiter,
			)
		}
	}
	return res
}

// toShellVariables create zsh statement where pmyOut's attributes are
// passed into shell variables
func (out *pmyOut) toShellVariables() string {
	res := ""
	res += fmt.Sprintf("local %v=$'%v';", shellCommandVariableName, utils.Escape(out.buildMainCommand(), "'"))
	res += fmt.Sprintf("local %v=$'%v';", shellBufferLeftVariableName, utils.Escape(out.bufferLeft, "'"))
	res += fmt.Sprintf("local %v=$'%v';", shellBufferRightVariableName, utils.Escape(out.bufferRight, "'"))
	res += fmt.Sprintf("local %v=$'%v';", shellFuzzyFinderCmdVariableName, utils.Escape(out.fuzzyFinderCmd, "'"))
	if out.allEmptyTag {
		res += fmt.Sprintf("local %v=$'%v';", shellTagAllEmptyVariableName, "empty")
	}

	for _, cg := range out.cmdGroups {
		res += fmt.Sprintf(
			"local %v=$'%v';",
			fmt.Sprintf(shellAfterVariableName, utils.EncodeTag(cg.Tag)),
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

// expandAllMagics expands all magic commnad in Stmt
// written in `%hoge` format
func (out *pmyOut) expandAllMagics() {
	for _, cg := range out.cmdGroups {
		if !strings.HasPrefix(cg.Stmt, "%") {
			continue
		}
		snippetBaseName := strings.Replace(cg.Stmt, "%", "", -1)
		snippetPath := fmt.Sprintf("%v/%v.txt", PmySnippetRoot, snippetBaseName)
		cg.Stmt = fmt.Sprintf(
			"cat %v | taggo -q '0:yellow' -d ' '",
			snippetPath,
		)
	}
	return
}

// expandAllParams expands all params that refer to regexp parameters
func (out *pmyOut) expandAllParams(paramMap map[string]string) {
	out.bufferLeft = expand(out.bufferLeft, paramMap)
	out.bufferRight = expand(out.bufferRight, paramMap)
	out.fuzzyFinderCmd = expand(out.fuzzyFinderCmd, paramMap)
	for _, cg := range out.cmdGroups {
		cg.Stmt = expand(cg.Stmt, paramMap)
	}
	return
}

func (out *pmyOut) toJSON() string {
	bytes, _ := json.Marshal(out)
	str := string(bytes)
	return str
}

// func (out *pmyOut) serialize() string {
//     return out.BufferLeft + delimiter + out.BufferRight + delimiter + out.Command
// }

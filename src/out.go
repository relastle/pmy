package pmy

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	utils "github.com/relastle/pmy/src/utils"
)

const (
	shellBufferLeftVariableName  = "__pmy_out_buffer_left"
	shellBufferRightVariableName = "__pmy_out_buffer_right"
	shellCommandVariableName     = "__pmy_out_command"
	shellSourcesVariableName     = "__pmy_out_sources"
	shellAfterVariableName       = "__pmy_out_%s_after"
	shellImmCmdVariableName      = "__pmy_out_imm_cmd"
	shellImmAfterCmdVariableName = "__pmy_out_imm_after_cmd"
)

type afterCmd struct {
	tag   string
	after string
}

// pmyOut represents Output of pmy against zsh routine.
// This struct has strings exported to shell, whose embedded
// variables are all expanded.
type pmyOut struct {
	bufferLeft  string
	bufferRight string
	cmdGroups   CmdGroups
	sources     string
	immCmd      string
	immAfterCmd string
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
	// expand all parameters
	out.expandAll(rule.paramMap)
	// get sources
	if immCmdGroup, ok := out.cmdGroups.getImmCmdGroup(); ok {
		out.immCmd = immCmdGroup.Stmt
		out.immAfterCmd = immCmdGroup.After
	} else {
		out.sources, _ = out.cmdGroups.GetSources()
	}
	// get sources
	return out
}

// Encode with base64 and then replace `/` and `+`
// into `a_a` and `b_b` respectively.
func encodeTag(tag string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(tag))
	sEnc = strings.ReplaceAll(sEnc, "/", "a_a")
	sEnc = strings.ReplaceAll(sEnc, "+", "b_b")
	sEnc = strings.ReplaceAll(sEnc, "=", "c_c")
	return sEnc
}

// toShellVariables create zsh statement where pmyOut's attributes are
// passed into shell variables
func (out *pmyOut) toShellVariables() string {
	res := ""
	res += fmt.Sprintf("%v=$'%v';", shellBufferLeftVariableName, utils.Escape(out.bufferLeft, "'"))
	res += fmt.Sprintf("%v=$'%v';", shellBufferRightVariableName, utils.Escape(out.bufferRight, "'"))
	res += fmt.Sprintf("%v=$'%v';", shellSourcesVariableName, utils.Escape(out.sources, "'"))
	res += fmt.Sprintf("%v=$'%v';", shellImmCmdVariableName, utils.Escape(out.immCmd, "'"))
	res += fmt.Sprintf("%v=$'%v';", shellImmAfterCmdVariableName, utils.Escape(out.immAfterCmd, "'"))
	for _, cg := range out.cmdGroups {
		res += fmt.Sprintf(
			"%v=$'%v';",
			fmt.Sprintf(shellAfterVariableName, encodeTag(cg.Tag)),
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

func (out *pmyOut) expandAll(paramMap map[string]string) {
	out.bufferLeft = expand(out.bufferLeft, paramMap)
	out.bufferRight = expand(out.bufferRight, paramMap)
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

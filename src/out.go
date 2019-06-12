package pmy

import (
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
)

type afterCmd struct {
	tag   string
	after string
}

// pmyOut represents Output of pmy against zsh routine.
// This struct has strings exported to shell, whose embedded
// variables are all expanded.
type pmyOut struct {
	BufferLeft  string `json:"bufferLeft"`
	BufferRight string `json:"bufferRight"`
	Sources     string `json:"source"`
	afterCmds   []*afterCmd
}

// newPmyOutFromRule create new pmyOut from rule
// which matches query and already has paramMap
func newPmyOutFromRule(rule *pmyRule) pmyOut {
	out := pmyOut{}
	// pass sources
	out.Sources, _ = rule.CmdGroups.GetSources()
	// pass resulting buffer informaiton
	out.BufferLeft = rule.BufferLeft
	out.BufferRight = rule.BufferRight
	// expand all parameters
	out.expandAll(rule.paramMap)
	// set after commnad
	out.afterCmds = []*afterCmd{}
	for _, cg := range rule.CmdGroups {
		ac := &afterCmd{
			tag:   cg.Tag,
			after: cg.After,
		}
		out.afterCmds = append(out.afterCmds, ac)
	}
	return out
}

// toShellVariables create zsh statement where pmyOut's attributes are
// passed into shell variables
func (out *pmyOut) toShellVariables() string {
	res := ""
	res += fmt.Sprintf("%v=$'%v';", shellBufferLeftVariableName, utils.Escape(out.BufferLeft, "'"))
	res += fmt.Sprintf("%v=$'%v';", shellBufferRightVariableName, utils.Escape(out.BufferRight, "'"))
	res += fmt.Sprintf("%v=$'%v';", shellSourcesVariableName, utils.Escape(out.Sources, "'"))
	for _, ac := range out.afterCmds {
		res += fmt.Sprintf(
			"%v=$'%v';",
			fmt.Sprintf(shellAfterVariableName, ac.tag),
			utils.Escape(ac.after, "'"),
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
	out.BufferLeft = expand(out.BufferLeft, paramMap)
	out.BufferRight = expand(out.BufferRight, paramMap)
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

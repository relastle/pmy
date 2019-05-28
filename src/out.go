package pmy

import (
	"encoding/json"
	"fmt"
	"strings"
)

// pmyOut represents Output of pmy against zsh routine.
// This struct has strings exported to shell, whose embedded
// variables are all expanded.
type pmyOut struct {
	BufferLeft  string `json:"bufferLeft"`
	BufferRight string `json:"bufferRight"`
	Command     string `json:"command"`
}

// newPmyOutFromRule create new pmyOut from rule
// which matches query and already has paramMap
func newPmyOutFromRule(rule *pmyRule) pmyOut {
	out := pmyOut{}
	out.Command = rule.Command
	out.BufferLeft = rule.BufferLeft
	out.BufferRight = rule.BufferRight
	out.expandAll(rule.paramMap)
	return out
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
	out.Command = expand(out.Command, paramMap)
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

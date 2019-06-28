package pmy

import (
	utils "github.com/relastle/pmy/src/utils"
)

// CmdGroup represents one command
// creating sources for fzf
type CmdGroup struct {
	Tag        string `json:"tag"`
	TagColor   string `json:"tagColor"`
	Stmt       string `json:"stmt"`
	After      string `json:"after"`
	tagAligned string
}

// CmdGroups is representing Slice of Cmd
// Resulting all commands will be one set of fzf-sources
type CmdGroups []*CmdGroup

// getMaxTagLen get max tag length given slice of command group
func (cgs CmdGroups) getMaxTagLen() int {
	maxLen := 0
	for _, cg := range cgs {
		if len(cg.Tag) > maxLen {
			maxLen = len(cg.Tag)
		}
	}
	return maxLen
}

// alignTag aligns tag by their lengths.
// This function set cg.`tagAligned`.
func (cgs CmdGroups) alignTag() {
	max := cgs.getMaxTagLen()
	for _, cg := range cgs {
		cg.tagAligned = cg.Tag + utils.MakeNString(max-len(cg.Tag), " ")
	}
	return
}

func (cgs CmdGroups) allEmpty() bool {
	for _, cg := range cgs {
		if cg.Tag != "" {
			return false
		}
	}
	return true
}

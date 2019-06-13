package pmy

import (
	"strings"

	pipeline "github.com/mattn/go-pipeline"
	shellwords "github.com/mattn/go-shellwords"
	utils "github.com/relastle/pmy/src/utils"
)

// CmdUnit represents a single command that does not have
// pipe operation or have shell environment embeddings
type CmdUnit struct {
	Stmt string   `json:"stmt"`
	Args []string `json:"args"`
}

// init initializes CmdUnit
func (cu *CmdUnit) init() error {
	// elms := strings.Split(cu.Stmt, " ")
	p := shellwords.NewParser()
	p.ParseEnv = true
	args, err := p.Parse(cu.Stmt)
	if err != nil {
		return err
	}
	cu.Args = args
	return nil
}

// CmdGroup represents one command
// creating sources for fzf
type CmdGroup struct {
	Tag      string    `json:"tag"`
	Stmt     string    `json:"stmt"`
	After    string    `json:"after"`
	CmdUnits []CmdUnit `json:"cmdUnits"`
	Lines    []string  `json:"lines"`
}

func (cg *CmdGroup) toPipeLine() []([]string) {
	res := []([]string){}
	for _, cu := range cg.CmdUnits {
		res = append(res, cu.Args)
	}
	return res
}

func (cg *CmdGroup) setLines(out []byte) {
	strList := strings.Split(strings.Trim(string(out), "\n"), "\n")
	cg.Lines = strList
	return
}

// init initializes CmdGroup by splitting
// stmt into piped commands, and then
// each command string will be converted into
// exec.Command executable structure.
func (cg *CmdGroup) init() {
	cg.CmdUnits = []CmdUnit{}
	pipedStmts := strings.Split(cg.Stmt, "|")
	for _, pipedStmt := range pipedStmts {
		cmdUnit := CmdUnit{
			Stmt: pipedStmt,
		}
		cmdUnit.init()
		cg.CmdUnits = append(cg.CmdUnits, cmdUnit)
	}
}

// CmdGroups is representing Slice of Cmd
// Resulting all commands will be one set of fzf-sources
type CmdGroups []*CmdGroup

func (cgs CmdGroups) getMaxTagLen() int {
	maxLen := 0
	for _, cg := range cgs {
		if len(cg.Tag) > maxLen {
			maxLen = len(cg.Tag)
		}
	}
	return maxLen
}

func (cgs CmdGroups) getImmCmdGroup() (*CmdGroup, bool) {
	if len(cgs) == 1 && cgs[0].Tag == "" {
		return cgs[0], true
	}
	return nil, false
}

func (cgs CmdGroups) organizeLines() string {
	resString := ""
	maxTagLen := cgs.getMaxTagLen()
	for _, cg := range cgs {
		spacedTag := cg.Tag + utils.MakeNString(3+maxTagLen-len(cg.Tag), " ")
		for _, line := range cg.Lines {
			resString += spacedTag + line + "\n"
		}
	}
	return resString
}

// GetSources get lines(fzf-sources) output by multiple command
// execution.
func (cgs CmdGroups) GetSources() (string, error) {
	// Prepare shellwords parser
	shellwords.ParseBacktick = true
	p := shellwords.NewParser()
	p.ParseEnv = true

	// Get output of each command group
	for _, cg := range cgs {
		cg.init()
		out, err := pipeline.Output(cg.toPipeLine()...)
		if err != nil {
			return "", nil
		}
		cg.setLines(out)
	}

	// Organize output
	return cgs.organizeLines(), nil
}

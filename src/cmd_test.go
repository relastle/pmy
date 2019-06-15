package pmy

import (
	"fmt"
	"testing"
)

func TestMultipleCommands(t *testing.T) {
	cg1 := CmdGroup{
		Tag:  "go",
		Stmt: "find . -maxdepth 2 | egrep .go$ | egrep -v test",
	}
	cg2 := CmdGroup{
		Tag:  "test go",
		Stmt: "find . -maxdepth 2 | egrep .go$ | egrep test",
	}
	cgs := CmdGroups{&cg1, &cg2}
	out, err := cgs.GetSources()
	if err != nil {
		t.Log(err)
	} else {
		fmt.Println(out)
	}
}

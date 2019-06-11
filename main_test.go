package main

import (
	"fmt"
	"testing"

	pmy "github.com/relastle/pmy/src"
)

func TestMultipleCommands(t *testing.T) {
	bufferLeft := "vim "
	bufferRight := ""

	cfgPath := "./resources/pmy_rules_test.json"

	outString := pmy.Run(
		cfgPath,
		pmy.Input{
			BufferLeft:  bufferLeft,
			BufferRight: bufferRight,
		},
	)
	fmt.Println(outString)
}

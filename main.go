package main

import (
	"flag"
	"fmt"

	"github.com/relastle/colorflag"
	pmy "github.com/relastle/pmy/src"
)

var (
	bufferLeft  string
	bufferRight string
)

func main() {
	flag.StringVar(&bufferLeft, "bufferLeft", "", "Current left buffer string of zsh prompt")
	flag.StringVar(&bufferRight, "bufferRight", "", "Current right buffer string of zsh prompt")

	colorflag.Parse([]*flag.FlagSet{})

	pmy.SetConfigs()

	outString := pmy.Run(
		pmy.PmyRulePath,
		pmy.Input{
			BufferLeft:  bufferLeft,
			BufferRight: bufferRight,
		},
	)
	fmt.Println(outString)
}

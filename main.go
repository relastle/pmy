package main

import (
	"flag"
	"fmt"

	pmy "github.com/relastle/pmy/src"
)

var (
	bufferLeft  string
	bufferRight string
)

func main() {
	flag.StringVar(&bufferLeft, "bufferLeft", "", "")
	flag.StringVar(&bufferRight, "bufferRight", "", "")
	flag.Parse()

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

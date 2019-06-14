package main

import (
	"flag"
	"fmt"

	pmy "github.com/relastle/pmy/src"
)

func main() {
	var bufferLeft string
	var bufferRight string
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

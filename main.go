package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/rakyll/statik/fs"
	_ "github.com/relastle/pmy/statik"

	"github.com/relastle/colorflag"
	pmy "github.com/relastle/pmy/src"
)

var (
	bufferLeft  string
	bufferRight string
)

func mainRoutine() {
	outString := pmy.Run(
		pmy.PmyRulePath,
		pmy.Input{
			BufferLeft:  bufferLeft,
			BufferRight: bufferRight,
		},
	)
	fmt.Println(outString)
}

func initCmdRoutine() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	f, err := statikFS.Open("/_pmy.zsh")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(bs))
}

func ruleListCmdRoutine() {

}

func snippetListCmdRoutine() {
}

func main() {
	flag.StringVar(&bufferLeft, "bufferLeft", "", "Current left buffer string of zsh prompt")
	flag.StringVar(&bufferRight, "bufferRight", "", "Current right buffer string of zsh prompt")

	// Subcommand for dumping init zsh script to stdout
	initFlagSet := flag.NewFlagSet("init", flag.ExitOnError)

	// Subcommandd for listing the all loaded rule json paths
	ruleListFlagSet := flag.NewFlagSet("rule-list", flag.ExitOnError)

	// Subcommand for listing the all loaded snippet json paths
	snippetsListFlagSet := flag.NewFlagSet("snippet-list", flag.ExitOnError)

	pmy.SetConfigs()

	subCommand := colorflag.Parse([]*flag.FlagSet{
		initFlagSet,
		ruleListFlagSet,
		snippetsListFlagSet,
	})

	switch subCommand {
	case "init":
		initCmdRoutine()
	case "rule-list":
		fmt.Println("rule-list")
	case "snippet-list":
		fmt.Println("snippet-list")
	default:
		mainRoutine()
	}
}

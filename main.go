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
	fmt.Printf("%s", string(bs))
}

func debug() {
}

func ruleListCmdRoutine() {
	pmy.SetConfigs()
	ruleFiles := pmy.GetAllRuleFiles()
	for _, ruleFile := range ruleFiles {
		fmt.Println(ruleFile.Path)
	}
}

func snippetListCmdRoutine() {
	pmy.SetConfigs()
	paths := pmy.GetAllSnippetJSONPaths()
	for _, path := range paths {
		fmt.Println(path)
	}
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

	// Subcommand for debuggin
	debugFlagSet := flag.NewFlagSet("debug", flag.ExitOnError)

	pmy.SetConfigs()

	subCommand := colorflag.Parse([]*flag.FlagSet{
		initFlagSet,
		ruleListFlagSet,
		snippetsListFlagSet,
		debugFlagSet,
	})

	switch subCommand {
	case "init":
		initCmdRoutine()
	case "rule-list":
		ruleListCmdRoutine()
	case "snippet-list":
		snippetListCmdRoutine()
	case "debug":
		debug()
	default:
		mainRoutine()
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/rakyll/statik/fs"
	_ "github.com/relastle/pmy/statik"
	"github.com/urfave/cli/v2"

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

func ruleListCmdRoutine() {
	ruleFiles := pmy.GetAllRuleFiles()
	for _, ruleFile := range ruleFiles {
		fmt.Println(ruleFile.Path)
	}
}

func snippetListCmdRoutine() {
	snippetFiles := pmy.GetAllSnippetFiles()
	for _, snippetFile := range snippetFiles {
		fmt.Println(snippetFile.Path)
	}
}

func main() {
	pmy.SetConfigs()

	start := time.Now()
	app := cli.NewApp()
	app.Version = "0.5.2"

	app.Commands = []*cli.Command{
		{
			Name:  "main",
			Usage: "Run main task of pmy. It dumps zsh script necessary to invoke fuzzy finder with appropriate source and edit current zsh line.",
			Action: func(c *cli.Context) error {
				mainRoutine()
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "buffer-left, l",
					Usage:       "Current `left` buffer of zsh line",
					Destination: &bufferLeft,
					Required:    true,
				},
				&cli.StringFlag{
					Name:        "buffer-right, r",
					Usage:       "Current `right` buffer of zsh line",
					Destination: &bufferRight,
					Required:    true,
				},
			},
		},
		{
			Name:  "init",
			Usage: "Dump zsh script to be sourced in order to activate pmy.",
			Action: func(c *cli.Context) error {
				initCmdRoutine()
				return nil
			},
		},
		{
			Name:  "rule-list",
			Usage: "List all rule files loaded.",
			Action: func(c *cli.Context) error {
				ruleListCmdRoutine()
				return nil
			},
		},
		{
			Name:  "snippet-list",
			Usage: "List all snippet files loaded.",
			Action: func(c *cli.Context) error {
				snippetListCmdRoutine()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	pmy.MeasureElapsedTime(start, "cli arg parse")

}

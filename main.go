package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	pmy "github.com/relastle/pmy/src"
)

var (
	bufferLeft  string
	bufferRight string
)

//go:embed _shell/_pmy.zsh
var initScript string

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
	fmt.Printf("%s", initScript)
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
	app.Version = "0.7.0"

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

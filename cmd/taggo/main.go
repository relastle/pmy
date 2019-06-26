package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var colorFuncMap = map[string](func(format string, a ...interface{}) string){
	"balck":   color.BlackString,
	"red":     color.RedString,
	"green":   color.BlueString,
	"yellow":  color.YellowString,
	"blue":    color.CyanString,
	"magenda": color.MagentaString,
	"cyan":    color.CyanString,
	"white":   color.WhiteString,
}

func main() {
	parse()
	color.NoColor = false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if index >= 0 {
			elms := strings.Split(text, delimiter)
			fmt.Printf(
				"%v%v%v\n",
				colorFuncMap[colorStr](elms[0]),
				delimiter,
				strings.Join(elms[1:len(elms)], delimiter),
			)
		} else {
			fmt.Printf(
				"%v%v%v\n",
				colorFuncMap[colorStr](tag),
				delimiter,
				text,
			)
		}
	}

	if scanner.Err() != nil {
	}
}

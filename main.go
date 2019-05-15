package main

import (
	"flag"
	"os"

	pmy "github.com/relastle/pmy/src"
)

const (
	// DefaultPmyConfigPath defiens the default config path
	// used when corresponding environment variable was not set.
	DefaultPmyConfigPath = ""
	// PmyConfigEnvVarName defiens the variable name
	// You should export this value
	PmyConfigEnvVarName = "PMY_CONFIG_PATH"
)

func main() {
	var bufferLeft string
	var bufferRight string
	flag.StringVar(&bufferLeft, "bufferLeft", "", "")
	flag.StringVar(&bufferRight, "bufferRight", "", "")
	flag.Parse()

	cfgPath, ok := os.LookupEnv(PmyConfigEnvVarName)
	if !ok {
		cfgPath = DefaultPmyConfigPath
	}
	pmy.Run(
		cfgPath,
		bufferLeft,
		bufferRight,
	)
}

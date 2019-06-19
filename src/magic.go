package pmy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getMagicSources(snippetBaseName string) string {
	snippetPath := fmt.Sprintf("%v/%v.txt", PmySnippetRoot, snippetBaseName)
	snippetFile, err := os.Open(snippetPath)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer snippetFile.Close()
	byteValue, _ := ioutil.ReadAll(snippetFile)
	magicOut := strings.Trim(string(byteValue), "\n")
	return magicOut
}

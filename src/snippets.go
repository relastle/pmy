package pmy

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mattn/go-zglob"
)

const pmySnippetSuffix = ".txt"
const _pmySnippetSuffix = "pmy_snippet.txt"

// SnippetFile represents one Snippet Json file
// information
type SnippetFile struct {
	Path     string
	Basename string
	Relpath  string
}

func (s SnippetFile) isApplicable(relpath string) bool {
	if s.Relpath == fmt.Sprintf(
		"%s%s",
		relpath,
		pmySnippetSuffix,
	) {
		return true
	}

	return s.Relpath == fmt.Sprintf(
		"%s_%s",
		relpath,
		_pmySnippetSuffix,
	)
}

// GetAllSnippetFiles get all pmy snippets txt paths
// configured by environment variable
func GetAllSnippetFiles() []*SnippetFile {
	snippetRoots := strings.Split(SnippetPath, ":")
	snippetRoots = append(snippetRoots, defaultSnippetPath)

	res := []*SnippetFile{}
	for _, snippetRoot := range snippetRoots {
		// expand environment variable
		snippetRoot = os.ExpandEnv(snippetRoot)
		if snippetRoot == "" {
			continue
		}
		globPattern := fmt.Sprintf(
			`%v/**/*%v`,
			snippetRoot,
			_pmySnippetSuffix,
		)
		matches, err := zglob.Glob(globPattern)
		if err != nil {
			panic(err)
		}
		for _, snippetPath := range matches {
			relpath, err := filepath.Rel(snippetRoot, snippetPath)
			if err != nil {
				log.Fatal(err)
				relpath = ""
			}
			res = append(
				res,
				&SnippetFile{
					Path:     snippetPath,
					Basename: path.Base(snippetPath),
					Relpath:  relpath,
				},
			)
		}

	}
	return res
}

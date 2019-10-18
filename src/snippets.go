package pmy

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mattn/go-zglob"
)

const pmySnippetSuffix = "pmy_snippet.txt"

// SnippetFile represents one Snippet Json file
// information
type SnippetFile struct {
	Path     string
	Basename string
	Relpath  string
}

func (s SnippetFile) isApplicable(relpath string) bool {
	return fmt.Sprintf(
		"%s_%s",
		relpath,
		pmySnippetSuffix,
	) == s.Relpath
}

// GetAllSnippetFiles get all pmy snippets txt paths
// configured by environment variable
func GetAllSnippetFiles() []*SnippetFile {
	snippetRoots := strings.Split(SnippetPath, ":")
	snippetRoots = append(snippetRoots, defaultSnippetPath)

	res := []*SnippetFile{}
	for _, snippetRoot := range snippetRoots {
		if snippetRoot == "" {
			continue
		}
		globPattern := fmt.Sprintf(
			`%v/**/*%v`,
			os.ExpandEnv(snippetRoot),
			pmySnippetSuffix,
		)
		matches, err := zglob.Glob(globPattern)
		if err != nil {
			panic(err)
		}
		for _, snippetPath := range matches {
			relpath, err := filepath.Rel(snippetRoot, snippetPath)
			if err != nil {
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

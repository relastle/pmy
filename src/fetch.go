package pmy

import (
	"log"
	"regexp"
)

type ruleFetcher interface {
	// GetCommand returns the command to be executed
	// given the
	fetch(
		rules pmyRules,
		bufferLeft string,
		bufferRight string,
	) (pmyRules, error)
}

// Mock of rruleFetcher
type ruleFetcherMock struct {
}

func (cg *ruleFetcherMock) fetch(
	rules pmyRules,
	bufferLeft string,
	bufferRight string,
) (pmyRules, error) {
	rule := pmyRule{
		RegexpLeft:  "make ",
		RegexpRight: "",
		Command:     "ls",
	}
	return pmyRules{rule}, nil
}

type ruleFetcherImpl struct {
}

func (cg *ruleFetcherImpl) fetch(
	rules pmyRules,
	bufferLeft string,
	bufferRight string,
) (pmyRules, error) {
	resRules := pmyRules{}
	for _, rule := range rules {
		if matched, err := regexp.Match(rule.RegexpLeft, []byte(bufferLeft)); matched {
			resRules = append(resRules, rule)
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return resRules, nil
}

package pmy

import "log"

type ruleFetcher interface {
	// GetCommand returns the command to be executed
	// given the
	fetch(
		rules Rules,
		bufferLeft string,
		bufferRight string,
	) (Rules, error)
}

// Mock of rruleFetcher
type ruleFetcherMock struct {
}

func (cg *ruleFetcherMock) fetch(
	rules Rules,
	bufferLeft string,
	bufferRight string,
) (Rules, error) {
	return nil, nil
}

type ruleFetcherImpl struct {
}

func (cg *ruleFetcherImpl) fetch(
	rules Rules,
	bufferLeft string,
	bufferRight string,
) (Rules, error) {
	resRules := Rules{}
	for _, rule := range rules {
		ok, err := rule.match(
			bufferLeft,
			bufferRight,
		)
		if err != nil {
			log.Fatal(err)
			continue
		}
		if ok {
			resRules = append(resRules, rule)
		}
	}
	return resRules, nil
}

package anypm

// Rule is a struct representing one rule
type Rule struct {
	regexpLeft  string `json:"regexpLeft"`
	regexpRight string `json:"regexpRight"`
	command     string `json:"command"`
}

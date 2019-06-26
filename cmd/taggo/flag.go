package main

import (
	"flag"
	"log"
)

var (
	tag       string
	colorStr  string
	delimiter string
	index     int
)

func checkColor() bool {
	ks := []string{}
	for k := range colorFuncMap {
		if colorStr == k {
			return true
		}
		ks = append(ks, k)
	}
	log.Fatalf("color must be any of %v\n", ks)
	return false
}

func parse() {
	flag.StringVar(&tag, "tag", "", "")
	flag.StringVar(&colorStr, "color", "black", "")
	flag.StringVar(&delimiter, "delimiter", "\t", "")
	flag.IntVar(&index, "index", -1, "")
	flag.Parse()
	// Validation
	checkColor()
}

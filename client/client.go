package main

import (
	"flag"
)

var (
	graphFile string
)

func main() {
	if !parseFlag() {
		flag.PrintDefaults()
		return
	}
}

func parseFlag() bool {
	flag.StringVar(
		&graphFile,
		"ungraph-source",
		"",
		"file that contains the edge list of the graph to construct",
	)

	flag.Parse()

	if graphFile == "" {
		return false
	}

	return true
}

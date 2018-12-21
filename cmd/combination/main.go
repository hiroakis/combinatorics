package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hiroakis/combinatorics"
)

var num = flag.Int("n", 3, "The number of choose")

func main() {
	flag.Parse()

	if flag.NArg() < *num {
		fmt.Println("The number of args should be greater than -n")
		os.Exit(1)
	}

	l := combinatorics.Unique(flag.Args())
	n := len(l)
	r := *num

	retval := combinatorics.Combination(l, n, r)
	for _, v := range retval {
		fmt.Printf("%+v\n", v)
	}

}

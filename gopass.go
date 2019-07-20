package main

import (
	"flag"
	"fmt"
)

var passLen, numPasses int

func init() {
	flag.IntVar(&passLen, "l", 12, "length of the password")
	flag.IntVar(&numPasses, "n", 1, "number of passwords")
	flag.Parse()
}

func main() {
	fmt.Println("Hello!")
	fmt.Println("Pass Len: ", passLen, " Num Passes: ", numPasses)
}

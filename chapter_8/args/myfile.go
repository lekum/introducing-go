package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()
	// Generate a number between 0 and max
	fmt.Println(rand.Intn(*maxp))
	args := flag.Args()
	if len(args) != 0 {
		fmt.Println("The non-flag arguments are", args)
	}
}

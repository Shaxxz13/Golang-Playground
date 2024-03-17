package main

import (
	"fmt"
	"runtime"
)

func main() {

	// Print OS Info
	os := runtime.GOOS
	fmt.Printf("Runtime OS: %v\n", os)

	args := os.Args

	if len(args) < 2 {

		fmt.Printf("Usage: ./main <argument> <argument>\n")
		os.Exit(10)
	}

	fmt.Printf("Total Arguments Entered: %v\n1st-Nth Entry: %v\n", args, args[1:])
}

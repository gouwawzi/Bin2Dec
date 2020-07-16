package main

import (
	"fmt"
	"os"
)
func main() {
	result, err := convert(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, result)
}

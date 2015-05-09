package main

import (
	"fmt"
	"os"
)

func main() {
	newLine := true
	parseFlags := true
	first := true
	for _, arg := range os.Args[1:] {
		if parseFlags && arg == "-n" {
			newLine = false
			continue
		}
		parseFlags = false
		if !first {
			fmt.Print(" ")
		}
		first = false
		fmt.Print(arg)
	}
	if newLine {
		fmt.Println()
	}
}

package main

import "os"
import "fmt"

func main() {
	var argsEnd int = 1
	var echoLen int = len(os.Args)

	// no newline flag
	if os.Args[1] == "-n" {
		argsEnd = 2
		echoLen = echoLen - 1
	}

	// echo each non-cmd/flag arg
	for i, a := range os.Args[argsEnd:] {
		c := ' '
		if i == argsEnd+echoLen {
			c = '\n'
		}
		fmt.Printf("%s%c", a, c)
	}
}

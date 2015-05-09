package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := "y"
	if len(os.Args) > 1 {
		msg = strings.Join(os.Args[1:], " ")
	}
	for {
		fmt.Println(msg)
	}
}

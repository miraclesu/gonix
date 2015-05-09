package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	flag.Parse()
	err := os.Remove(flag.Arg(0))

	if err != nil {
		fmt.Println(err)
	}
}

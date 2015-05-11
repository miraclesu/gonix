package main

import "fmt"
import "flag"
import "path"

func main() {
	flag.Parse()
	if len(flag.Args()) == 1 {
		fmt.Println(path.Base(flag.Arg(0)))
	} else if len(flag.Args()) == 2 {
		base   := path.Base(flag.Arg(0))
		suffix := flag.Arg(1)
		if base[len(base)-len(suffix):] == suffix {
			fmt.Println(base[:len(base)-len(suffix)])
		} else {
			fmt.Println(base)
		}
	}
}

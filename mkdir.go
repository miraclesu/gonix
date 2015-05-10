package main

import "os"
import "flag"
import "log"

func main() {
	pFlag := flag.Bool("p",false,"create parent directories if needed")
	flag.Parse()
	if len(flag.Args())>0 {
		for i:=0; i<len(flag.Args()); i++ {
			if (*pFlag) {
				err := os.MkdirAll(flag.Arg(i), 0644)
			} else {
				err := os.Mkdir(flag.Arg(i), 0644)
			}
			if err != nil {log.Fatal(err)}
		}
	}
}

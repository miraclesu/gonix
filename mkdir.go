package main

import "os"
import "log"

func main() {
	if len(os.Args)>0 {
		for i:=1; i<len(os.Args); i++ {
			err := os.Mkdir(os.Args[i], 0644)
			if err != nil {log.Fatal(err)}
		}
	}
}

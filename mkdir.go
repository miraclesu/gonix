package main

import "os"

func main() {
	if len(os.Args)>0 {
		for i:=1; i<len(os.Args); i++ {
			os.Mkdir(os.Args[i], 0644)
		}
	}
}

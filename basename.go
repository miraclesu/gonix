package main

import "fmt"
import "os"
import "path"

func main() {
	if len(os.Args) > 1 {
		fmt.Println(path.Base(os.Args[1]))
	}
}

package main

import "fmt"
import "os"
import "path"

func main() {
	if len(os.Args) > 1 {
		pathname := os.Args[1]
		//If the given path ends in a slash, delete the last character of it (the slash)
		if pathname[len(pathname)-1] == '/' {
			pathname = pathname[:1]
		}
		base := path.Base(pathname)
		fmt.Println(base)
	}
}

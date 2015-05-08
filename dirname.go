package main

import "fmt"
import "os"
import "path"

func main() {
    if(len(os.Args)>1) {
        dir := path.Dir(os.Args[1])
        fmt.Println(dir)
    }
}

package main

import "os"
import "flag"

func main() {
    flag.Parse()
    os.Mkdir(flag.Arg(0),0644)
}

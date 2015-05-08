package main

import "os"
import "fmt"
import "io/ioutil"

func main() {
    if(len(os.Args)==1) {
        bytes, _ := ioutil.ReadAll(os.Stdin)
        fmt.Printf("%s",string(bytes))
    } else {
        for i:=1; i<len(os.Args); i++ {
            bytes, _ := ioutil.ReadFile(os.Args[i])
            fmt.Printf("%s",string(bytes))
        }
    }
}

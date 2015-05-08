package main

import "os"
import "fmt"
import "io/ioutil"
import "flag"

func main() {
    flagAppend := flag.Bool("a",false,"append to file")
    flag.Parse()
    bytes, _ := ioutil.ReadAll(os.Stdin);
    for i:=0; i<len(flag.Args()); i++ {
        if(*flagAppend) {
            f,_ := os.OpenFile(flag.Args()[i], os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
            f.Write(bytes)
            f.Close()
        } else {
            f,_ := os.OpenFile(flag.Args()[i], os.O_WRONLY | os.O_CREATE, 0644)
            f.Write(bytes)
            f.Close()
        }
    }
    fmt.Printf("%s",string(bytes))
}

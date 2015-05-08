package main

import "os"
import "fmt"

func main() {
    if(len(os.Args)==1) {
        for ;; {
            fmt.Println("y")
        }
    } else {
        for ;; {
            for i:=1; i<len(os.Args); i++ {
                fmt.Printf("%s ",os.Args[i])
            }
            fmt.Printf("\n")
        }
    }
}

package main

import "os"
import "time"
import "strconv"

func main() {
    if(len(os.Args)==2) {
        n,_ := strconv.Atoi(os.Args[1])
        time.Sleep(time.Second * time.Duration(n))
    }
}

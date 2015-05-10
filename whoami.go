package main

import "os/user"
import "fmt"
import "log"

func main() {
    u, err := user.Current()
    if err != nul {log.Fatal(err)}
    fmt.Println(u.Username)
}

package main

import "fmt"
import "log"
import "os"

func main() {
	wd, err := os.Getwd()
	if err != nil {log.Fatal(err)}
	fmt.Println(wd)
}

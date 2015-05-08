package main

import "fmt"
import "os"

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
}

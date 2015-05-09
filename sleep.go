package main

import "os"
import "log"
import "time"
import "strconv"

func main() {
	if len(os.Args) == 2 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {log.Fatal(err)}
		time.Sleep(time.Second * time.Duration(n))
	}
}

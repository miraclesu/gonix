package main

import "os"
import "time"

func main() {
	filename := os.Args[1]
	_, err := os.Stat(filename)
	if err == nil {
		now := time.Now()
		os.Chtimes(filename, now, now)
	} else {
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		f.Close()
	}
}

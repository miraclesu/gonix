package main

import "os"
import "flag"
import "time"
import "log"

func main() {
	cFlag := flag.Bool("c",false,"do not create file")
	flag.Parse()
	if len(flag.Args())>0 {
		for i:=0; i<len(flag.Args()); i++ {
			filename := flag.Arg(i)
			_, err := os.Stat(filename)
			if err == nil {
				now := time.Now()
				os.Chtimes(filename, now, now)
			} else {
				if !(*cFlag) {
					f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
					f.Close()
					if err != nil {log.Fatal(err)}
				}
			}
		}
	}
}

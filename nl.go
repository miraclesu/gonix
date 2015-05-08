package main

import "os"
import "fmt"
import "io/ioutil"
import "strings"
import "flag"

func main() {
	bFlag := flag.String("b", "t", "style")
	flag.Parse()
	if len(flag.Args()) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		lines := strings.Split(string(bytes), "\n")
		linecount := 0
		for i := 0; i < len(lines)-1; i++ {
			if *bFlag == "t" {
				if len(strings.TrimSpace(lines[i])) > 0 {
					linecount++
					fmt.Printf("%6d  %s\n", linecount, lines[i])
				} else {
					fmt.Println(lines[i])
				}
			} else {
				linecount++
				fmt.Printf("%6d  %s\n", linecount, lines[i])
			}
		}
	} else if len(flag.Args()) > 0 {
		linecount := 0
		for j := 0; j < len(flag.Args()); j++ {
			bytes, _ := ioutil.ReadFile(flag.Arg(j))
			lines := strings.Split(string(bytes), "\n")
			for i := 0; i < len(lines)-1; i++ {
				if *bFlag == "t" {
					if len(strings.TrimSpace(lines[i])) > 0 {
						linecount++
						fmt.Printf("%6d  %s\n", linecount, lines[i])
					} else {
						fmt.Println(lines[i])
					}
				} else {
					linecount++
					fmt.Printf("%6d  %s\n", linecount, lines[i])
				}
			}
		}
	}
}

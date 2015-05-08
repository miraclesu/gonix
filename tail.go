package main

import "os"
import "fmt"
import "io/ioutil"
import "strings"
import "flag"

func main() {
	n := flag.Int("n", 10, "number of lines")
	flag.Parse()
	if len(flag.Args()) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		lines := strings.Split(string(bytes), "\n")
		var end int
		//Make sure we don't print more lines than there are
		if *n > len(lines)-1 {
			end = len(lines) - 1
		} else {
			end = *n
		}
		for i := (len(lines) - 1 - (end)); i < len(lines)-1; i++ {
			fmt.Println(lines[i])
		}
	} else if len(flag.Args()) > 0 {
		for j := 0; j < len(flag.Args()); j++ {
			bytes, _ := ioutil.ReadFile(flag.Arg(j))
			lines := strings.Split(string(bytes), "\n")
			var end int
			//Make sure we don't print more lines than there are
			if *n > len(lines)-1 {
				end = len(lines) - 1
			} else {
				end = *n
			}
			//Only show a title for each file if there are multiple files
			if len(flag.Args()) > 1 {
				fmt.Printf("==> %s <==\n", flag.Arg(j))
			}
			for i := (len(lines) - 1 - (end)); i < len(lines)-1; i++ {
				fmt.Println(lines[i])
			}
			if j < len(flag.Args())-1 {
				fmt.Printf("\n")
			}
		}
	}
}

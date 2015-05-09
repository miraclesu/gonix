package main

import "os"
import "fmt"
import "log"
import "io/ioutil"

func main() {
	if len(os.Args) == 1 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if (err != nil) {log.Fatal(err)}
		fmt.Printf("%s", string(bytes))
	} else {
		for i := 1; i < len(os.Args); i++ {
			bytes, err := ioutil.ReadFile(os.Args[i])
			if (err != nil) {log.Fatal(err)}
			fmt.Printf("%s", string(bytes))
		}
	}
}

package main

import "os"
import "fmt"
import "log"
import flag "github.com/ogier/pflag"
import "bufio"

func main() {
	n := flag.IntP("lines", "n", 10, "number of lines")
	q := flag.BoolP("quiet", "q", false, "do not print filenames")
	flag.Parse()
	if len(flag.Args()) == 0 {
		reader := bufio.NewReader(os.Stdin)
		for linecounter := 0; ; {
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				break
			}
			if linecounter >= *n {
				break
			}
			if !isPrefix {
				linecounter++
			}
			fmt.Printf("%s\n", string(line))
		}
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			file, err := os.Open(flag.Arg(i))
			if err != nil {
				log.Fatal(err)
			}
			reader := bufio.NewReader(file)
			if !(*q) && len(flag.Args()) > 1 {
				if i > 0 {
					//Seperate the files with a blank line
					fmt.Println()
				}
				//Print the name of each file
				fmt.Printf("==> %s <==\n", flag.Arg(i))
			}
			for linecounter := 0; ; {
				line, isPrefix, err := reader.ReadLine()
				if err != nil {
					break
				}
				if linecounter >= *n {
					break
				}
				if !isPrefix {
					//Don't increment linecounter if the line isn't finished
					linecounter++
				}
				fmt.Printf("%s\n", string(line))
			}
		}
	}
}

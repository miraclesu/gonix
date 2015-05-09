package main

import "os"
import "fmt"
import "log"
import "flag"
import "bufio"

func main() {
	n := flag.Int("n", 10, "number of lines")
	flag.Parse()
	if len(flag.Args()) == 0 {
		//Using stdin isn't working for some reason so piping to head won't work.
                //I need to fix this section
		reader := bufio.NewReader(os.Stdin)
		for linecounter := 0; ; {
			line, isPrefix, err := reader.ReadLine()
			if err != nil {break}
			if linecounter >= *n {break}
			if !isPrefix {break}
			fmt.Printf("%s\n", string(line))
		}
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			file, err := os.Open(flag.Arg(i))
			if err != nil {log.Fatal(err)}
			reader := bufio.NewReader(file)
			for linecounter := 0; ; {
				line, isPrefix, err := reader.ReadLine()
				if err != nil {break}
				if linecounter >= *n {break}
				if !isPrefix {linecounter++} //Don't increment linecounter if the line isn't finished
				fmt.Printf("%s\n", string(line))
			}
		}
	}
}

package main

import (
	"os"
	"strconv"
)
import "fmt"

func main() {
	// TODO: Also handle floats.
	start, inc, end := 1, 1, 0
	nargs := len(os.Args)
	switch nargs {
	case 4:
		inc = parseInt(os.Args[2])
		fallthrough
	case 3:
		start = parseInt(os.Args[1])
		fallthrough
	case 2:
		end = parseInt(os.Args[nargs-1])
	default:
		msg := "missing operand"
		if nargs > 4 {
			msg = fmt.Sprintf("extra operand '%s'", os.Args[4])
		}
		fmt.Println("seq:", msg)
		os.Exit(1)
	}

	for i := start; i <= end; i += inc {
		fmt.Println(i)
	}
}

func parseInt(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("seq:", err)
		os.Exit(1)
	}
	return i
}

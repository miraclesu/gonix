package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// chose digest function
	var newHash = md5.New

	// no args, read from stdin
	if len(os.Args) == 1 {
		h := newHash()
		_, err := io.Copy(h, os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%x\t%s\n", h.Sum(nil), "-")
		return
	}

	// md5 string digest
	var s = flag.String("s", "", "print the checksum of the given string")
	flag.Parse()
	sVal := *s
	if sVal != "" {
		fmt.Printf("%x\t\"%s\"\n", newHash().Sum([]byte(sVal)), sVal)
		return
	}

	// md5 files & stdin digest
	for _, filename := range flag.Args() {
		var (
			f   = os.Stdin
			h   = newHash()
			err error
		)
		if filename != "-" {
			f, err = os.Open(filename)
		}
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		_, err = io.Copy(h, f)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("%x\t%s\n", h.Sum(nil), filename)
	}

}

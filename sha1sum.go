package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// choose a digest function
	var newHash = sha1.New

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

	// string digest
	var s = flag.String("s", "", "print the checksum of the given string")
	flag.Parse()
	sVal := *s
	if sVal != "" {
		fmt.Printf("%x\t\"%s\"\n", newHash().Sum([]byte(sVal)), sVal)
		return
	}

	// files & stdin digests
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

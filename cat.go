package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			log.Fatal(err)
		}
	} else {
		for i := 1; i < len(os.Args); i++ {
			if err := catFile(os.Args[i]); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func catFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(os.Stdout, file)
	return err
}

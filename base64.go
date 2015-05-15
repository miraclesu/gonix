/*
	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License version 3 as
	published by the Free Software Foundation.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func readData(reader io.Reader) ([]byte, error) {
	return ioutil.ReadAll(reader)
}

func isAlpha(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

func readAndHandle(reader io.Reader, flagDecode *bool, flagIgnore *bool, flagWrap *int) {
	src, err := readData(reader)
	checkError(err)
	var toHandle []byte
	if *flagIgnore {
		//It seems that the effect of "base64 -i" in *nix
		//is not filter the non-alphabet charater.
		//This flag cannot work as the original *nix command flag.
		for i := 0; i < len(src); i++ {
			if isAlpha(src[i]) {
				toHandle = append(toHandle, src[i])
			}
		}
	} else {
		toHandle = src
	}
	if *flagDecode {
		decoded, err := base64Decode(toHandle)
		checkError(err)
		fmt.Printf("%s", string(decoded))
	} else {
		encoded := base64Encode(toHandle)
		wrapPrint(encoded, *flagWrap)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func wrapPrint(output []byte, wrap int) {
	if wrap == 0 {
		fmt.Printf("%s\n", string(output))
		return
	}

	length := len(output)
	if length <= wrap {
		fmt.Printf("%s\n", string(output))
		return
	}

	index, end := 0, 0
	for index < length {
		end += wrap
		if end > length {
			end = length
		}
		fmt.Printf("%s\n", string(output[index:end]))
		index += wrap
	}
}

func main() {
	flagDecode := flag.Bool("d", false, "Decode the data")
	flagIgnore := flag.Bool("i", false, "When decoding, ignore non-alphabet characters")
	flagWrap := flag.Int("w", 76, "Wrap encoded lines after COLS character (default 76). Use 0 to disable line wrapping.")

	flag.Parse()
	if *flagWrap < 0 {
		log.Fatalf("invalid wrap size: %d", *flagWrap)
	}

	if len(flag.Args()) == 0 {
		readAndHandle(os.Stdin, flagDecode, flagIgnore, flagWrap)
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			file, err := os.Open(flag.Args()[i])
			checkError(err)
			readAndHandle(file, flagDecode, flagIgnore, flagWrap)
		}
	}
}

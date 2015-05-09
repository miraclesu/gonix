package main 
import (
	"unicode"
	"log"
	"os"
	"io/ioutil"
	"io"
	"flag"
	"fmt"
	"encoding/base64"
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

func readAndHandle(reader io.Reader, flagDecode *bool, flagIgnore *bool) {
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
		fmt.Printf("%s\n", string(decoded))
	} else {
		encoded := base64Encode(toHandle)
		fmt.Printf("%s\n", string(encoded))
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flagDecode := flag.Bool("d", false, "Decode the data")
	flagIgnore := flag.Bool("i", false, "When decoding, ignore non-alphabet characters")
	//TODO: -w
	flag.Parse()
	if len(flag.Args()) == 0 {
		readAndHandle(os.Stdin, flagDecode, flagIgnore)	
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			file, err := os.Open(flag.Args()[i])
			checkError(err)
			readAndHandle(file, flagDecode, flagIgnore)
		}
	}
}

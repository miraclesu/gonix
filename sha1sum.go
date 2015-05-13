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

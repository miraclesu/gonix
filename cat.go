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

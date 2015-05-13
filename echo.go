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
	"fmt"
	"os"
)

func main() {
	newLine := true
	parseFlags := true
	first := true
	for _, arg := range os.Args[1:] {
		if parseFlags && arg == "-n" {
			newLine = false
			continue
		}
		parseFlags = false
		if !first {
			fmt.Print(" ")
		}
		first = false
		fmt.Print(arg)
	}
	if newLine {
		fmt.Println()
	}
}

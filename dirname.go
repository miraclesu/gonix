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

import "fmt"
import "os"
import "path"

func main() {
	if len(os.Args) > 1 {
		pathname := os.Args[1]
		//If the given path ends in a slash, delete the last character of it (the slash)
		if len(pathname) > 1 && pathname[len(pathname)-1] == '/' {
			pathname = pathname[:len(pathname)-1]
		}
		fmt.Println(path.Dir(pathname))
	}
}
